package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type User struct {
	Id       int `gorm:"primaryKey"`
	Name     string
	Username string
	Email    string
	Password string
	Gender   string
}

type LoginReq struct {
	Name     string
	Password string
}

func NewUser() (User, error) {
	var user User

	fmt.Print("Masukkan Nama : ")
	_, err := fmt.Scanln(&user.Name)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan Username : ")
	_, err = fmt.Scanln(&user.Username)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan email : ")
	_, err = fmt.Scanln(&user.Email)
	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan password : ")
	_, err = fmt.Scanln(&user.Password)

	if err != nil {
		return user, err
	}

	fmt.Print("Masukkan gender : ")
	_, err = fmt.Scanln(&user.Gender)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (usr User) UserToString() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s", usr.Name, usr.Username, usr.Email, usr.Password, usr.Gender)
}

func NewLoginReq() (LoginReq, error) {
	var loginReq LoginReq

	fmt.Print("Masukkan Name : ")
	_, err := fmt.Scanln(&loginReq.Name)
	if err != nil {
		return loginReq, err
	}

	fmt.Print("Masukkan Name : ")
	_, err = fmt.Scanln(&loginReq.Password)
	if err != nil {
		return loginReq, err
	}

	return loginReq, err
}

func Login() bool {
	var req, err = NewLoginReq()
	if err != nil {
		return false
	}

	db, err := ConnectDB()
	if err != nil {
		return false
	}

	var user User
	if err := db.Where("name = ?", req.Name).First(&user).Error; err != nil {
		return false
	}

	passReq := md5.Sum([]byte(req.Password))
	passString := hex.EncodeToString(passReq[:])

	if passString == user.Password {
		return true
	}
	return false
}

func Profile(name string) (User, error) {
	var user User

	db, err := ConnectDB()
	if err != nil {
		return user, err
	}

	if err := db.Where("name = ?", name).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
