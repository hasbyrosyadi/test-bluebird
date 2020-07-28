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
	role      string
	IsLogin   int
}
