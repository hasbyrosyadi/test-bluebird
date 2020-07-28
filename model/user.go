package model

import (
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
