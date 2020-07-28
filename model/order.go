package model

import (
	"time"
)

type Order struct {
	Id          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	CartId      int    `json:"cart_id"`
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	ProductName string `json:"product_name"`
	Address     string `json:"address"`
	Quantity    int    `json:"quantity"`
}
