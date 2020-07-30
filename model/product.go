package model

import (
	"errors"
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

func AddProduct(product *ReqProduct) (*Product, error) {
	if product.ProductName == "" || product.Price == 0 || product.Stock == 0 {
		return nil, errors.New("Missing Parameter")
	}

	addProduct := &Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	return addProduct, nil
}

func (p *Product) UpdateProduct(product *EditProduct) (bool, error) {
	if product.ProductName == "" || product.Price == 0 || product.Stock == 0 {
		return false, errors.New("Missing Parameter")
	}

	var isUpdated bool
	if p.ProductName != product.ProductName {
		p.ProductName = product.ProductName
		isUpdated = true
	}

	if p.Price != product.Price {
		p.Price = product.Price
		isUpdated = true
	}

	if p.Stock != product.Stock {
		p.Stock = product.Stock
		isUpdated = true
	}

	return isUpdated, nil
}
