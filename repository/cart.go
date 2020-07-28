package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	db *gorm.DB
}

func NewCart(d *gorm.DB) CartRepository {
	return &Cart{db: d}
}

type CartRepository interface {
	GetAllProduct() ([]model.Cart, error)
}

func (c *Cart) GetAllProduct() ([]model.Cart, error) {
	cart := []model.Cart{}
	err := c.db.Find(cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
