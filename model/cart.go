package model

import (
	"time"
)

type Cart struct {
	Id          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	OrderId     int    `json:"order_id"`
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}
