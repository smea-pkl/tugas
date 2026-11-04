package main

import (
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
