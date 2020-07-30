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
	// menampilkan seluruh product
	product, err := p.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Server Error")
	}
	return product, nil
}

func (p *Product) AddProduct(email string, product *model.ReqProduct) error {

	// get user berdasarkan email
	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	// validasi user sebagai admin
	err = user.IsAdmin()
	if err != nil {
		return err
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return err
	}

	// validasi product input
	newProduct, err := model.AddProduct(product)
	if err != nil {
		return err
	}

	// insert product
	err = p.ProductRepository.InsertProduct(newProduct)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}

func (p *Product) EditProduct(email string, product *model.EditProduct) (*model.Product, error) {

	// get user berdasarkan email
	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi user sebagai admin
	err = user.IsAdmin()
	if err != nil {
		return nil, err
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return nil, err
	}

	// validasi kalo product tersebut terdaftar
	dataProduct, err := p.ProductRepository.GetProductById(product.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if dataProduct.Id == 0 {
		return nil, errors.New("Product not found")
	}

	// validasi serta update kolom yang mau diubah
	isUpdated, err := dataProduct.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	// update product
	if isUpdated {
		err = p.ProductRepository.UpdateProduct(dataProduct)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}
	}

	return dataProduct, nil
}

func (p *Product) DeleteProduct(email string, id int) error {

	// get user berdasarkan email
	user, err := p.UserRepository.GetUser(email)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	// validasi user sebagai admin
	err = user.IsAdmin()
	if err != nil {
		return err
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return err
	}

	// validasi kalo product tersebut terdaftar
	dataProduct, err := p.ProductRepository.GetProductById(id)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	if dataProduct.Id == 0 {
		return errors.New("Product not Found")
	}

	// delete product secara soft delete
	err = p.ProductRepository.DeleteProduct(id)
	if err != nil {
		return errors.New("Internal Server Error")
	}

	return nil
}
