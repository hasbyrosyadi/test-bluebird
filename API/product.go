package api

import (
	"bluebird/model"
	"bluebird/usecase"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Product struct {
	ProductUsecase usecase.ProductUsecase
}

func (p *Product) GetAll(w http.ResponseWriter, r *http.Request) {
	product, err := p.ProductUsecase.GetAllProduct()
	if err != nil {
		errorResponse := ErrorMessage(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(product)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (p *Product) InsertProduct(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	if email == "" {
		errorResponse := ErrorClient(errors.New("Unauthorized"))
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	if err := r.ParseForm(); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	var reqProduct model.ReqProduct
	if err := schema.NewDecoder().Decode(&reqProduct, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	err := p.ProductUsecase.AddProduct(email, &reqProduct)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(nil)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return

}

func NewProductHandler(router *mux.Router, usecase usecase.ProductUsecase) {
	handler := &Product{
		ProductUsecase: usecase,
	}

	router.HandleFunc("/products", handler.GetAll).Methods("GET")
	router.HandleFunc("/add_products", handler.InsertProduct).Methods("POST")

}
