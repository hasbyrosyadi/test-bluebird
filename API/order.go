package api

import (
	"bluebird/model"
	"bluebird/usecase"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Order struct {
	OrderUsecase usecase.OrderUsecase
}

func (o *Order) PostOrder(w http.ResponseWriter, r *http.Request) {
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

	var order model.PostOrder
	if err := schema.NewDecoder().Decode(&order, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	resp, err := o.OrderUsecase.PostOrder(email, order.CartId)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(resp)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (o *Order) PaidOrder(w http.ResponseWriter, r *http.Request) {
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

	var order model.PaidOrder
	if err := schema.NewDecoder().Decode(&order, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	resp, err := o.OrderUsecase.PaidOrder(email, order.OrderId)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(resp)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (o *Order) HistoryOrder(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	if email == "" {
		errorResponse := ErrorClient(errors.New("Unauthorized"))
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	resp, err := o.OrderUsecase.HistoryOrder(email)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(resp)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func NewOrderHandler(router *mux.Router, usecase usecase.OrderUsecase) {
	handler := &Order{
		OrderUsecase: usecase,
	}

	router.HandleFunc("/post_order", handler.PostOrder).Methods("POST")
	router.HandleFunc("/paid_order", handler.PaidOrder).Methods("POST")
	router.HandleFunc("/history_order", handler.HistoryOrder).Methods("GET")
}
