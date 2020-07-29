package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"errors"
)

type ProductUsecase interface {
	GetAllProduct() ([]model.Product, error)
	AddProduct(email string, product *model.ReqProduct) error
	EditProduct(email string, product *model.EditProduct) (*model.Product, error)
	DeleteProduct(email string, id int) error
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

	if product.ProductName == "" || product.Price == 0 || product.Stock == 0 {
		return errors.New("Missing Parameter")
	}

	addProduct := &model.Product{
		ProductName: product.ProductName,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	err = p.ProductRepository.InsertProduct(addProduct)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (p *Product) EditProduct(email string, product *model.EditProduct) (*model.Product, error) {

	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if user.Role != "ADMIN" || user.IsLogin != true || user.Id == 0 {
		return nil, errors.New("Access Denied")
	}

	dataProduct, err := p.ProductRepository.GetProductById(product.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if product.ProductName == "" || product.Price == 0 || product.Stock == 0 {
		return nil, errors.New("Missing Parameter")
	}

	var isUpdated bool
	if dataProduct.ProductName != product.ProductName {
		dataProduct.ProductName = product.ProductName
		isUpdated = true
	}

	if dataProduct.Price != product.Price {
		dataProduct.Price = product.Price
		isUpdated = true
	}

	if dataProduct.Stock != product.Stock {
		dataProduct.Stock = product.Stock
		isUpdated = true
	}

	if isUpdated {
		err = p.ProductRepository.UpdateProduct(dataProduct)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}
	}

	return dataProduct, nil
}

func (p *Product) DeleteProduct(email string, id int) error {

	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if user.Role != "ADMIN" || user.IsLogin != true || user.Id == 0 {
		return errors.New("Access Denied")
	}

	dataProduct, err := p.ProductRepository.GetProductById(id)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if dataProduct.Id == 0 {
		return errors.New("Product not Found")
	}

	err = p.ProductRepository.DeleteProduct(id)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}
