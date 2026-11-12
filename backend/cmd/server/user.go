package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"time"

	"connectrpc.com/connect"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id       string `gorm:"primaryKey"`
	Username string
	Email    string
	Password string
}

type userServiceImpl struct{}

// Register implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Register(ctx context.Context, req *connect.Request[usersv1.RegisterRequest]) (*connect.Response[usersv1.RegisterResponse], error) {
	pay := req.Msg

	passByte := md5.Sum([]byte(pay.Password))
	passHash := hex.EncodeToString(passByte[:])

	user := User{
		Username: pay.Username,
		Email:    pay.Email,
		Password: passHash,
	}

	db, err := ConnectDB()
	if err != nil {
		res := usersv1.RegisterResponse{
			Msg: "Gagal connect database",
		}
		return connect.NewResponse(&res), err
	}

	err = db.Save(&user).Error
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

var jwtKey = []byte("owi")

// Login implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Login(ctx context.Context, req *connect.Request[usersv1.LoginRequest]) (*connect.Response[usersv1.LoginResponse], error) {
	pay := req.Msg
	db, err := ConnectDB()
	if err != nil {
		res := usersv1.LoginResponse{
			Token: "",
			Msg:   "Gagal connect database",
		}
		return connect.NewResponse(&res), err
	}

	var user User
	if err := db.Where("username = ?", pay.Username).First(&user).Error; err != nil {
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
		return connect.NewResponse(&res), err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(1 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		res := usersv1.LoginResponse{
			Token: "",
			Msg:   "Tidak bisa membuat token",
		}
		return connect.NewResponse(&res), err
	}

	res := usersv1.LoginResponse{
		Token: tokenString,
		Msg:   "Berhasil Login",
	}
	return connect.NewResponse(&res), nil
}
