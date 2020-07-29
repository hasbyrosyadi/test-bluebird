package model

import (
	"time"
)

type Cart struct {
	Id            int `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	OrderId       *int
	UserId        int
	ConfigProduct string
	TotalPayment  float64
	Quantity      int
}

type AddToCart struct {
	ConfigProduct string `schema:"config_product"`
}

type DetailConfig struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type ConfigProduct struct {
	Id       int
	Price    float64
	Quantity int
}

type RespAddToCart struct {
	Id            int            `json:"id"`
	TotalPayment  float64        `json:"total_payment"`
	TotalQuantity int            `json:"total_quantity"`
	ConfigProduct []DetailConfig `json:"list"`
}
