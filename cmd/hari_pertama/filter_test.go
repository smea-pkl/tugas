package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserInput(t *testing.T) {
	t.Run("User input sutarman", func(t *testing.T) {
		usr := User{
			Name:     "Sutarman",
			Username: "Sutarman",
			Email:    "sutarman@gmail.com",
			Password: "sutarman",
			Gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input andika", func(t *testing.T) {
		usr := User{
			Name:     "Andika",
			Username: "Andika",
			Email:    "andika@gmail.com",
			Password: "andika",
			Gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input fitri dan gender p", func(t *testing.T) {
		usr := User{
			Name:     "Fitri",
			Username: "Fitri",
			Email:    "fitri@gmail.com",
			Password: "fitri",
			Gender:   "p",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input fitri dan gender l", func(t *testing.T) {
		usr := User{
			Name:     "Fitri",
			Username: "Fitri",
			Email:    "fitri@gmail.com",
			Password: "fitri",
			Gender:   "l",
		}

		check, err := FilterUser(usr)
		assert.True(t, check)
		assert.Nil(t, err)
	})

	t.Run("User input gender selain p, P, l, L", func(t *testing.T) {
		usr := User{
			Name:     "Rizki",
			Username: "rizki",
			Email:    "rizki@gmail.com",
			Password: "rizki",
			Gender:   "T",
		}

		check, err := FilterUser(usr)
		assert.False(t, check)
		assert.Nil(t, err)
	})
}
