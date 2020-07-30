package model

import (
	"errors"
	"time"
)

type User struct {
	Id        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Email     string
	Password  string
	Address   string
	Role      string
	IsLogin   bool
}

type Register struct {
	Name            string `schema:"name"`
	Email           string `schema:"email"`
	Password        string `schema:"password"`
	ConfirmPassword string `schema:"confirm_password"`
	Address         string `schema:"address"`
}

type Login struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type Logout struct {
	Email string `schema:"email"`
}

func Regis(regis *Register) *User {
	newUser := &User{
		Name:     regis.Name,
		Email:    regis.Email,
		Password: regis.Password,
		Address:  regis.Address,
		Role:     "USER",
	}
	return newUser
}

func (u *User) IsAdmin() error {
	if u.Role != "ADMIN" {
		return errors.New("Access Denied")
	}
	return nil
}

func (u *User) IsLoginUser() error {
	if u.IsLogin != true || u.Id == 0 {
		return errors.New("Access Denied")
	}
	return nil
}
