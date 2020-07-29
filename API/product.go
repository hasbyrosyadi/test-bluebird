package api

import (
	"bluebird/model"
	"bluebird/usecase"
	"errors"
	"net/http"
	"strconv"

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

func (p *Product) EditProduct(w http.ResponseWriter, r *http.Request) {
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

	var editProduct model.EditProduct
	if err := schema.NewDecoder().Decode(&editProduct, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	product, err := p.ProductUsecase.EditProduct(email, &editProduct)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(product)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	idProduct, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		errorResponse := ErrorClient(errors.New("Missing Parameter"))
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	err = p.ProductUsecase.DeleteProduct(email, idProduct)
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
	router.HandleFunc("/edit_products", handler.EditProduct).Methods("POST")
	router.HandleFunc("/delete_products/{id}", handler.DeleteProduct).Methods("DELETE")
}
