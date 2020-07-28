package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"errors"
)

type ProductUsecase interface {
	GetAllProduct() ([]model.Product, error)
	AddProduct(email string, product *model.ReqProduct) error
}

type Product struct {
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
}

func NewProduct(p repository.ProductRepository, u repository.UserRepository) ProductUsecase {
	return &Product{p, u}
}

func (p *Product) GetAllProduct() ([]model.Product, error) {
	product, err := p.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Server Error")
	}
	return product, nil
}

func (p *Product) AddProduct(email string, product *model.ReqProduct) error {

	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Role != "ADMIN" || user.IsLogin != true || user.Id == 0 {
		return errors.New("Access Denied")
	}

	if product.ProductName == "" || product.Stock == 0 {
		return errors.New("Missing Parameter")
	}

	addProduct := &model.Product{
		ProductName: product.ProductName,
		Stock:       product.Stock,
	}

	err = p.ProductRepository.InsertProduct(addProduct)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}
