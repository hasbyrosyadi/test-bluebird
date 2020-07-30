package usecase

import (
	"bluebird/model"
	"bluebird/repository"
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

	// get user berdasarkan email
	user, err := c.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi user login dan sudah terdaftar
	if err = user.IsLoginUser(); err != nil {
		return nil, err
	}

	// verifikasi input
	if cart.ConfigProduct == "" {
		return nil, errors.New("Missing Parameter")
	}

	// get all product
	allProduct, err := c.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	totalPayment, totalQuantity, respConfigCart, err := model.ConfigProductOrder(cart.ConfigProduct, allProduct)
	if err != nil {
		return nil, err
	}

	addCart := &model.Cart{
		UserId:        user.Id,
		ConfigProduct: cart.ConfigProduct,
		TotalPayment:  totalPayment,
		Quantity:      totalQuantity,
	}

	// validasi cart yang masih aktif
	activeCart, err := c.CartRepository.GetActiveCart(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	var dataCart *model.Cart
	if activeCart.Id != 0 {
		// update cart
		activeCart.ConfigProduct = cart.ConfigProduct
		activeCart.TotalPayment = totalPayment
		activeCart.Quantity = totalQuantity
		dataCart, err = c.CartRepository.UpdateCart(activeCart)
		if err != nil {
			return nil, errors.New("Internal Server Error")
		}
	} else {

		// insert cart
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
