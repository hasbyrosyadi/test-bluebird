package model

import (
	"time"
)

type HistoryOrder struct {
	Id            int `gorm:"primary_key"`
	CreatedAt     time.Time
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
