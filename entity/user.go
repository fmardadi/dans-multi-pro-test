package entity

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) HashPassword(password string) error {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hashPass)
	fmt.Println(user.Password)
	return nil
}

func (user *User) CheckPassword(inputPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		return err
	}
	return nil
}
