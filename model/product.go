package model

import (
	"time"
)

type Product struct {
	Id          int        `gorm:"primary_key";json:"id"`
	CreatedAt   time.Time  `json:"_"`
	UpdatedAt   time.Time  `json:"_"`
	DeletedAt   *time.Time `json:"_"`
	ProductName string     `json:"product_name"`
	Stock       int        `json:"stock"`
}

type ReqProduct struct {
	ProductName string `schema:"product_name"`
	Stock       int    `schema:"stock"`
}
