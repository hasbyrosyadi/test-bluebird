package model

import (
	"time"
)

type Order struct {
	Id          int        `gorm:"primary_key";json:"id"`
	CreatedAt   time.Time  `json:"_"`
	UpdatedAt   time.Time  `json:"_"`
	DeletedAt   *time.Time `json:"_"`
	CartId      int        `json:"cart_id"`
	UserId      int        `json:"user_id"`
	Name        string     `json:"name"`
	ProductName string     `json:"product_name"`
	Address     string     `json:"address"`
	Quantity    int        `json:"quantity"`
}
