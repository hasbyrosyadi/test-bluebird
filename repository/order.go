package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type Order struct {
	db *gorm.DB
}

func NewOrder(d *gorm.DB) OrderRepository {
	return &Product{db: d}
}

type OrderRepository interface {
	GetAllProduct() ([]model.Product, error)
}

func (o *Order) GetAllProduct() ([]model.Order, error) {
	order := []model.Order{}
	err := o.db.Find(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
