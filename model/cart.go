package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Cart struct {
	Id            int `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	OrderId       *int
	UserId        int
	ConfigProduct string
	TotalPayment  float64
	Quantity      int
}

type AddToCart struct {
	ConfigProduct string `schema:"config_product"`
}

type DetailConfig struct {
	Id          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type ConfigProduct struct {
	Id       int
	Price    float64
	Quantity int
}

type RespAddToCart struct {
	Id            int            `json:"id"`
	TotalPayment  float64        `json:"total_payment"`
	TotalQuantity int            `json:"total_quantity"`
	ConfigProduct []DetailConfig `json:"list"`
}

func ConfigProductOrder(cart string, listProduct []Product) (float64, int, []DetailConfig, error) {

	productList := make(map[int]Product)

	var (
		totalPayment  float64
		totalQuantity int
	)

	// marshal config cart product
	var confiqCart []ConfigProduct
	if err := json.Unmarshal([]byte(cart), &confiqCart); err != nil {
		return 0, 0, nil, errors.New("Type Config is wrong")
	}

	var respConfigCart []DetailConfig

	// masukan product ke dalam list product
	for _, prod := range listProduct {
		productList[prod.Id] = prod
	}

	for _, data := range confiqCart {
		// validasi kesediaan stock
		if productList[data.Id].Stock == 0 || productList[data.Id].Stock < data.Quantity {
			return 0, 0, nil, errors.New("Product out of stock")
		}

		// validasi perubahan harga
		if productList[data.Id].Price != data.Price {
			return 0, 0, nil, errors.New("Price was Changed")
		}

		respConfigCart = append(respConfigCart, DetailConfig{
			Id:          productList[data.Id].Id,
			ProductName: productList[data.Id].ProductName,
			Price:       data.Price,
			Quantity:    data.Quantity,
		})

		totalPayment += data.Price * float64(data.Quantity)
		totalQuantity += data.Quantity
	}

	return totalPayment, totalQuantity, respConfigCart, nil
}
