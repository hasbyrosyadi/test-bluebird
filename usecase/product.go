package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"errors"
)

type ProductUsecase interface {
	GetAllProduct() ([]model.Product, error)
}

type Product struct {
	ProductRepository repository.ProductRepository
}

func NewProduct(p repository.ProductRepository) ProductUsecase {
	return &Product{p}
}

func (p *Product) GetAllProduct() ([]model.Product, error) {
	product, err := p.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Server Error")
	}
	return product, nil
}
