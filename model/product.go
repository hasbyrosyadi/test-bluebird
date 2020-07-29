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
	Price       float64    `json:"Price"`
	Stock       int        `json:"stock"`
}

type ReqProduct struct {
	ProductName string  `schema:"product_name"`
	Price       float64 `schema:"price"`
	Stock       int     `schema:"stock"`
}

type EditProduct struct {
	Id          int     `schema:"id"`
	ProductName string  `schema:"product_name"`
	Price       float64 `schema:"price"`
	Stock       int     `schema:"stock"`
}
