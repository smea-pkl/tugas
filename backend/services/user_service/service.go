package user_service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"pkl/backend/models"
	"pkl/backend/pkgs/san_config"
	"strings"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type userServiceImpl struct {
	db     *gorm.DB
	jwtKey []byte
}

// GetUser implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) GetUser(ctx context.Context, req *connect.Request[usersv1.GetUserRequest]) (*connect.Response[usersv1.GetUserResponse], error) {
	pay := req.Header().Get("Authorization")

	if strings.TrimSpace(pay) == "" {
		return nil, fmt.Errorf("Token kosong")
	}

	pay = strings.Replace(pay, "Bearer ", "", -1)

	cc := CustomClaims{}

	if err := cc.Verify(pay); err != nil {
		return nil, err
	}

	user := models.User{}
	if err := u.db.Where("username = ?", cc.Username).First(&user).Error; err != nil {
		return nil, err
	}

	return connect.NewResponse(&usersv1.GetUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}), nil
}

type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func (c *CustomClaims) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: c.Username,
	})

	cfg, err := san_config.RenderYamlConfig()
	if err != nil {
		return "", err
	}

	return token.SignedString([]byte(cfg.JwtKey))
}

func (c *CustomClaims) Verify(token string) error {
	cfg, err := san_config.RenderYamlConfig()
	if err != nil {
		return err
	}

	tkn, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (any, error) { return []byte(cfg.JwtKey), nil }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return err
	}

	if claims, ok := tkn.Claims.(*CustomClaims); ok && tkn.Valid {
		c.Username = claims.Username
		return nil
	}

	return fmt.Errorf("Invalid token")
}

// Register implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Register(ctx context.Context, req *connect.Request[usersv1.RegisterRequest]) (*connect.Response[usersv1.RegisterResponse], error) {
	pay := req.Msg

	passByte := md5.Sum([]byte(pay.Password))
	passHash := hex.EncodeToString(passByte[:])

	user := models.User{
		Username: pay.Username,
		Email:    pay.Email,
		Password: passHash,
	}

	err := u.db.Save(&user).Error
	if err != nil {
		res := usersv1.RegisterResponse{
			Msg: "Gagal menambahkan user",
		}
		return connect.NewResponse(&res), err
	}

	res := usersv1.RegisterResponse{
		Msg: "Berhasil register",
	}
	return connect.NewResponse(&res), nil
}

// Login implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Login(ctx context.Context, req *connect.Request[usersv1.LoginRequest]) (*connect.Response[usersv1.LoginResponse], error) {
	pay := req.Msg

	var user models.User
	if err := u.db.Where("username = ?", pay.Username).First(&user).Error; err != nil {
		res := usersv1.LoginResponse{
			Token: "",
			Msg:   "Username atau Password salah",
		}
		return connect.NewResponse(&res), err
	}

	passReq := md5.Sum([]byte(pay.Password))
	passHash := hex.EncodeToString(passReq[:])

	if passHash != user.Password {
		res := usersv1.LoginResponse{
			Token: "",
			Msg:   "Username atau Password salah",
		}
		return connect.NewResponse(&res), errors.New("Password Salah")
	}

	cc := CustomClaims{
		Username: user.Username,
	}

	token, err := cc.Generate()
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&usersv1.LoginResponse{Token: token, Msg: "Berhasil Login"}), nil
}
