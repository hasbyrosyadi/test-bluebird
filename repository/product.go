package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type Product struct {
	db *gorm.DB
}

func NewProduct(d *gorm.DB) ProductRepository {
	return &Product{db: d}
}

type ProductRepository interface {
	GetAllProduct() ([]model.Product, error)
}

func (p *Product) GetAllProduct() ([]model.Product, error) {
	product := []model.Product{}
	err := p.db.Find(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
