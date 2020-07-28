package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type Order struct {
	db *gorm.DB
}

func NewOrder(d *gorm.DB) OrderRepository {
	return &Order{db: d}
}

type OrderRepository interface {
	GetAllProduct() ([]model.Order, error)
}

func (o *Order) GetAllProduct() ([]model.Order, error) {
	order := []model.Order{}
	err := o.db.Find(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
