package api

import (
	"bluebird/model"
	"bluebird/usecase"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Cart struct {
	CartUsecase usecase.CartUsecase
}

func (c *Cart) AddToCart(w http.ResponseWriter, r *http.Request) {
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

	var addtoCart model.AddToCart
	if err := schema.NewDecoder().Decode(&addtoCart, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	resp, err := c.CartUsecase.AddToCart(email, &addtoCart)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(resp)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func NewCartHandler(router *mux.Router, usecase usecase.CartUsecase) {
	handler := &Cart{
		CartUsecase: usecase,
	}

	router.HandleFunc("/add_to_cart", handler.AddToCart).Methods("POST")
}
