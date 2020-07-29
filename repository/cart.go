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
	GetActiveCart(userId int) (*model.Cart, error)
	GetCartById(id int) (*model.Cart, error)
	InsertCart(cart *model.Cart) (*model.Cart, error)
	UpdateCart(cart *model.Cart) (*model.Cart, error)
}

func (c *Cart) GetActiveCart(userId int) (*model.Cart, error) {
	cart := &model.Cart{}
	err := c.db.Where("user_id = ? AND order_id IS NULL", userId).First(&cart).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return cart, nil
}

func (c *Cart) GetCartById(id int) (*model.Cart, error) {
	cart := &model.Cart{}
	err := c.db.Where("id = ?", id).First(&cart).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return cart, nil
}

func (c *Cart) InsertCart(cart *model.Cart) (*model.Cart, error) {
	if err := c.db.Create(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (c *Cart) UpdateCart(cart *model.Cart) (*model.Cart, error) {
	if err := c.db.Save(cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
