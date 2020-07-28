package api

import (
	"bluebird/model"
	"bluebird/usecase"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type User struct {
	UserUsecase usecase.UserUsecase
}

func (u *User) Register(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	var reqRegister model.Register
	if err := schema.NewDecoder().Decode(&reqRegister, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	err := u.UserUsecase.Register(&reqRegister)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(nil)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	var reqLogin model.Login
	if err := schema.NewDecoder().Decode(&reqLogin, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	err := u.UserUsecase.Login(&reqLogin)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(nil)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	var reqLogout model.Logout
	if err := schema.NewDecoder().Decode(&reqLogout, r.PostForm); err != nil {
		resp := ErrorMessage(err)
		HttpResponseJson(w, resp, resp.HTTPStatusCode)
		return
	}

	err := u.UserUsecase.Logout(&reqLogout)
	if err != nil {
		errorResponse := ErrorClient(err)
		HttpResponseJson(w, errorResponse, errorResponse.HTTPStatusCode)
		return
	}

	response := Success(nil)
	HttpResponseJson(w, response, response.HTTPStatusCode)
	return
}

func NewUserHandler(router *mux.Router, usecase usecase.UserUsecase) {
	handler := &User{
		UserUsecase: usecase,
	}

	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/logout", handler.Logout).Methods("POST")
}
