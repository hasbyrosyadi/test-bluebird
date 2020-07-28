package model

import (
	"time"
)

type Product struct {
	Id          int `gorm:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	ProductName string `json:"product_name"`
	Stock       int    `json:"stock"`
}
