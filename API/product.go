package api

import (
	"bluebird/usecase"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	ProductUsecase usecase.ProductUsecase
}

func (p *Product) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	product, err := p.ProductUsecase.GetAllProduct()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	response := Success(product)
	respBody, err := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.HTTPStatusCode)
	w.Write(respBody)
	return

}

func NewProductHandler(router *mux.Router, usecase usecase.ProductUsecase) {
	handler := &Product{
		ProductUsecase: usecase,
	}

	router.HandleFunc("/products", handler.GetAllProduct).Methods("GET")
}
