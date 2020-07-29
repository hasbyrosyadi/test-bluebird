package usecase

import (
	"bluebird/model"
	"bluebird/repository"
	"encoding/json"
	"errors"
)

type OrderUsecase interface {
	PostOrder(email string, cartId int) (*model.RespOrder, error)
	PaidOrder(email string, orderId int) (*model.RespOrder, error)
	HistoryOrder(email string) ([]model.RespOrder, error)
}

type Order struct {
	OrderRepository        repository.OrderRepository
	HistoryOrderRepository repository.HistoryOrderRepository
	CartRepository         repository.CartRepository
	ProductRepository      repository.ProductRepository
	UserRepository         repository.UserRepository
}

func NewOrder(o repository.OrderRepository, h repository.HistoryOrderRepository, c repository.CartRepository, p repository.ProductRepository, u repository.UserRepository) OrderUsecase {
	return &Order{o, h, c, p, u}
}

func (o *Order) PostOrder(email string, cartId int) (*model.RespOrder, error) {

	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if user.IsLogin != true || user.Id == 0 {
		return nil, errors.New("Access Denied")
	}

	if cartId == 0 {
		return nil, errors.New("Missing Parameter")
	}

	activeOrder, err := o.OrderRepository.GetActiveOrder(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if activeOrder.Id != 0 {
		return nil, errors.New("Please complete your bill")
	}

	cart, err := o.CartRepository.GetCartById(cartId)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if cart.Id == 0 {
		return nil, errors.New("Cart not Found")
	}

	var confiqCart []model.ConfigProduct
	if err := json.Unmarshal([]byte(cart.ConfigProduct), &confiqCart); err != nil {
		return nil, errors.New("Type Config is wrong")
	}

	var respConfigCart []model.DetailConfig

	for _, data := range confiqCart {

		product, err := o.ProductRepository.GetProductById(data.Id)
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
	}

	postOrder := &model.Order{
		CartId:        cart.Id,
		UserId:        cart.UserId,
		Name:          user.Name,
		ConfigProduct: cart.ConfigProduct,
		Address:       user.Address,
		TotalPayment:  cart.TotalPayment,
		Quantity:      cart.Quantity,
		Status:        "UNPAID",
	}

	dataOrder, err := o.OrderRepository.CreateOrder(postOrder)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	resp := &model.RespOrder{
		Id:            dataOrder.Id,
		CartId:        dataOrder.CartId,
		UserId:        dataOrder.UserId,
		Name:          dataOrder.Name,
		ConfigProduct: respConfigCart,
		Address:       dataOrder.Address,
		TotalPayment:  dataOrder.TotalPayment,
		Quantity:      dataOrder.Quantity,
		Status:        dataOrder.Status,
	}

	return resp, nil
}

func (o *Order) PaidOrder(email string, orderId int) (*model.RespOrder, error) {

	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if user.IsLogin != true || user.Id == 0 {
		return nil, errors.New("Access Denied")
	}

	if orderId == 0 {
		return nil, errors.New("Missing Parameter")
	}

	order, err := o.OrderRepository.GetOrderUnpaidById(orderId)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if order.Id == 0 {
		return nil, errors.New("Order not found")
	}

	var confiqCart []model.ConfigProduct
	if err := json.Unmarshal([]byte(order.ConfigProduct), &confiqCart); err != nil {
		return nil, errors.New("Type Config is wrong")
	}

	var respConfigCart []model.DetailConfig

	for _, data := range confiqCart {

		product, err := o.ProductRepository.GetProductById(data.Id)
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
	}

	dataOrder, err := o.OrderRepository.PaidOrder(order, confiqCart)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	resp := &model.RespOrder{
		Id:            dataOrder.Id,
		CartId:        dataOrder.CartId,
		UserId:        dataOrder.UserId,
		Name:          dataOrder.Name,
		ConfigProduct: respConfigCart,
		Address:       dataOrder.Address,
		TotalPayment:  dataOrder.TotalPayment,
		Quantity:      dataOrder.Quantity,
		Status:        dataOrder.Status,
	}

	return resp, nil
}

func (o *Order) HistoryOrder(email string) ([]model.RespOrder, error) {

	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if user.IsLogin != true || user.Id == 0 {
		return nil, errors.New("Access Denied")
	}

	historys, err := o.HistoryOrderRepository.GetAllHistory(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	var respHistory []model.RespOrder

	for _, history := range historys {

		var confiqCart []model.ConfigProduct
		if err := json.Unmarshal([]byte(history.ConfigProduct), &confiqCart); err != nil {
			return nil, errors.New("Type Config is wrong")
		}

		var respConfigCart []model.DetailConfig

		for _, config := range confiqCart {
			product, err := o.ProductRepository.GetProductById(config.Id)
			if err != nil {
				return nil, errors.New("Internal Server Error")
			}

			respConfigCart = append(respConfigCart, model.DetailConfig{
				Id:          product.Id,
				ProductName: product.ProductName,
				Price:       config.Price,
				Quantity:    config.Quantity,
			})
		}

		respHistory = append(respHistory, model.RespOrder{
			Id:            history.Id,
			CartId:        history.CartId,
			UserId:        history.UserId,
			Name:          history.Name,
			ConfigProduct: respConfigCart,
			Address:       history.Address,
			TotalPayment:  history.TotalPayment,
			Quantity:      history.Quantity,
			Status:        history.Status,
		})
	}

	return respHistory, nil
}
