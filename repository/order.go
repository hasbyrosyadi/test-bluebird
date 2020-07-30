package repository

import (
	"bluebird/model"

	"github.com/jinzhu/gorm"
)

type Order struct {
	db *gorm.DB
}

func NewOrder(d *gorm.DB) OrderRepository {
	return &Order{db: d}
}

type OrderRepository interface {
	GetActiveOrder(userId int) (*model.Order, error)
	GetOrderUnpaidById(id int) (*model.Order, error)
	PaidOrder(order *model.Order, confiqCart []model.DetailConfig) (*model.Order, error)
	CreateOrder(order *model.Order) (*model.Order, error)
}

func (o *Order) GetActiveOrder(userId int) (*model.Order, error) {
	order := &model.Order{}
	err := o.db.Where("user_id = ? AND status = ?", userId, "UNPAID").First(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return order, nil
}

func (o *Order) GetOrderUnpaidById(id int) (*model.Order, error) {
	order := &model.Order{}
	err := o.db.Where("id = ? AND status = ?", id, "UNPAID").First(&order).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return order, nil
}

func (o *Order) CreateOrder(order *model.Order) (*model.Order, error) {
	var err error
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err = tx.Create(&order).Error; err != nil {
		return nil, err
	}

	postHistoryOrder := &model.HistoryOrder{
		CartId:        order.Id,
		UserId:        order.UserId,
		Name:          order.Name,
		ConfigProduct: order.ConfigProduct,
		Address:       order.Address,
		TotalPayment:  order.TotalPayment,
		Quantity:      order.Quantity,
		Status:        "UNPAID",
	}

	if err = tx.Create(&postHistoryOrder).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) PaidOrder(order *model.Order, confiqCart []model.DetailConfig) (*model.Order, error) {
	var err error
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	order.Status = "PAID"
	if err = tx.Save(&order).Error; err != nil {
		return nil, err
	}

	for _, config := range confiqCart {
		product := &model.Product{}
		if err = tx.Where("id = ?", config.Id).First(product).Error; err != nil {
			return nil, err
		}
		stock := product.Stock
		product.Stock = stock - config.Quantity

		if err = tx.Save(&product).Error; err != nil {
			return nil, err
		}
	}

	cart := &model.Cart{}
	if err = tx.Model(cart).Update("order_id", order.Id).Error; err != nil {
		return nil, err
	}

	postHistoryOrder := &model.HistoryOrder{
		CartId:        order.Id,
		UserId:        order.UserId,
		Name:          order.Name,
		ConfigProduct: order.ConfigProduct,
		Address:       order.Address,
		TotalPayment:  order.TotalPayment,
		Quantity:      order.Quantity,
		Status:        "PAID",
	}

	if err = tx.Create(&postHistoryOrder).Error; err != nil {
		return nil, err
	}

	return order, nil
}
