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

	// get user berdasarkan email
	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return nil, err
	}

	// verifikasi input
	if cartId == 0 {
		return nil, errors.New("Missing Parameter")
	}

	// verifikasi order apakah sudah di submit atau belum
	activeOrder, err := o.OrderRepository.GetActiveOrder(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if activeOrder.Id != 0 {
		return nil, errors.New("Please complete your bill")
	}

	// verifikasi cart apakah sudah di order atau belum
	cart, err := o.CartRepository.GetCartById(cartId)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi apabila bukan user tersebut yang melakukan transaksi
	if cart.Id == 0 || cart.UserId != user.Id {
		return nil, errors.New("Cart not Found")
	}

	// get all product
	allProduct, err := o.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// configurasi cart
	_, _, respConfigCart, err := model.ConfigProductOrder(cart.ConfigProduct, allProduct)
	if err != nil {
		return nil, err
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

	// create order dan menambah history order
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

	// get user berdasarkan email
	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return nil, err
	}

	// verifikasi input
	if orderId == 0 {
		return nil, errors.New("Missing Parameter")
	}

	// get order user yang masih berstatus UNPAID
	order, err := o.OrderRepository.GetOrderUnpaidById(orderId)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi apabila bukan user tersebut yang melakukan transaksi
	if order.Id == 0 || order.UserId != user.Id {
		return nil, errors.New("Order not found")
	}

	// get all product
	allProduct, err := o.ProductRepository.GetAllProduct()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// configurasi cart
	_, _, respConfigCart, err := model.ConfigProductOrder(order.ConfigProduct, allProduct)
	if err != nil {
		return nil, err
	}

	// ubah status jadi paid, pengurangan stock di product dan menambah history order
	dataOrder, err := o.OrderRepository.PaidOrder(order, respConfigCart)
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

	// get user berdasarkan email
	user, err := o.UserRepository.GetUser(email)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	// validasi user login dan sudah terdaftar
	err = user.IsLoginUser()
	if err != nil {
		return nil, err
	}

	// get all history user login
	historys, err := o.HistoryOrderRepository.GetAllHistory(user.Id)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	var respHistory []model.RespOrder

	for _, history := range historys {

		// unmarshal config product
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
