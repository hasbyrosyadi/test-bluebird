package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"encoding/json"
	"errors"
)

type CartUsecase interface {
	AddToCart(email string, cart *model.AddToCart) (*model.RespAddToCart, error)
}

type Cart struct {
	CartRepository    repository.CartRepository
	ProductRepository repository.ProductRepository
	UserRepository    repository.UserRepository
}

func NewCart(c repository.CartRepository, p repository.ProductRepository, u repository.UserRepository) CartUsecase {
	return &Cart{c, p, u}
}

func (c *Cart) AddToCart(email string, cart *model.AddToCart) (*model.RespAddToCart, error) {

	user, err := c.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if user.IsLogin != true || user.Id == 0 {
		return nil, errors.New("Access Denied")
	}

	if cart.ConfigProduct == "" {
		return nil, errors.New("Missing Parameter")
	}

	var (
		totalPayment  float64
		totalQuantity int
	)

	var confiqCart []model.ConfigProduct
	if err := json.Unmarshal([]byte(cart.ConfigProduct), &confiqCart); err != nil {
		return nil, errors.New("Type Config is wrong")
	}

	var respConfigCart []model.DetailConfig

	for _, data := range confiqCart {

		product, err := c.ProductRepository.GetProductById(data.Id)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}

		if product.Stock == 0 || product.Stock < data.Quantity {
			return nil, errors.New("Product out of stock")
		}

		if product.Price != data.Price {
			return nil, errors.New("Price was Changed")
		}

		respConfigCart = append(respConfigCart, model.DetailConfig{
			Id:          product.Id,
			ProductName: product.ProductName,
			Price:       data.Price,
			Quantity:    data.Quantity,
		})

		totalPayment += data.Price
		totalQuantity += data.Quantity
	}

	addCart := &model.Cart{
		UserId:        user.Id,
		ConfigProduct: cart.ConfigProduct,
		TotalPayment:  totalPayment,
		Quantity:      totalQuantity,
	}

	activeCart, err := c.CartRepository.GetActiveCart(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	var dataCart *model.Cart
	if activeCart.Id != 0 {
		dataCart, err = c.CartRepository.UpdateCart(addCart)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}
	} else {
		dataCart, err = c.CartRepository.InsertCart(addCart)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}
	}

	resp := &model.RespAddToCart{
		Id:            dataCart.Id,
		TotalPayment:  totalPayment,
		TotalQuantity: totalQuantity,
		ConfigProduct: respConfigCart,
	}

	return resp, nil
}
