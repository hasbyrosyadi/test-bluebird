package model

import (
	"time"
)

type Cart struct {
	Id          int        `gorm:"primary_key";json:"id"`
	CreatedAt   time.Time  `json:"_"`
	UpdatedAt   time.Time  `json:"_"`
	DeletedAt   *time.Time `json:"_"`
	OrderId     int        `json:"order_id"`
	ProductId   int        `json:"product_id"`
	ProductName string     `json:"product_name"`
	Quantity    int        `json:"quantity"`
}
