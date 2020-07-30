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
	GetProductById(id int) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	InsertProduct(product *model.Product) error
	DeleteProduct(id int) error
}

func (p *Product) GetAllProduct() ([]model.Product, error) {
	product := []model.Product{}
	err := p.db.Find(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return product, nil
}

func (p *Product) GetProductById(id int) (*model.Product, error) {
	product := &model.Product{}
	err := p.db.Where("id = ?", id).First(&product).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return product, nil
}

func (p *Product) UpdateProduct(product *model.Product) error {
	if err := p.db.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) InsertProduct(product *model.Product) error {
	if err := p.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (p *Product) DeleteProduct(id int) error {
	if err := p.db.Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
		return err
	}
	return nil
}
