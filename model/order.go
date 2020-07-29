package model

import (
	"time"
)

type Order struct {
	Id            int `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	CartId        int
	UserId        int
	Name          string
	ConfigProduct string
	Address       string
	TotalPayment  float64
	Quantity      int
	Status        string
}

type PostOrder struct {
	CartId int `schema:"cart_id"`
}

type PaidOrder struct {
	OrderId int `schema:"order_id"`
}

type RespOrder struct {
	Id            int            `json:"id"`
	CartId        int            `json:"cart_id"`
	UserId        int            `json:"user_id"`
	Name          string         `json:"name"`
	ConfigProduct []DetailConfig `json:"list"`
	Address       string         `json:"address"`
	TotalPayment  float64        `json:"total_payment"`
	Quantity      int            `json:"quantity"`
	Status        string         `json:"status"`
}
