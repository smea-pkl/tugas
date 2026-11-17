package user_service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"pkl/backend/models"
	"strings"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type userServiceImpl struct {
	db     *gorm.DB
	jwtKey []byte
}

var secretKey = []byte("rahasia")

type CustomClaims struct {
	Username string
	jwt.RegisteredClaims
}

func (c *CustomClaims) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Username: c.Username,
	})

	return token.SignedString(secretKey)
}

func (c *CustomClaims) Verify(token string) error {
	tkn, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(t *jwt.Token) (any, error) { return secretKey, nil }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return err
	}

	if claims, ok := tkn.Claims.(*CustomClaims); ok && tkn.Valid {
		c.Username = claims.Username
		return nil
	}

	return fmt.Errorf("Invalid token")
}

// GetProfile implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) GetProfile(ctx context.Context, req *connect.Request[usersv1.GetProfileRequest]) (*connect.Response[usersv1.GetProfileResponse], error) {
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

	return connect.NewResponse(&usersv1.GetProfileResponse{
		Username: user.Username,
		Email:    user.Email,
	}), nil
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
