package main

import (
	api "bluebird/API"
	"bluebird/repository"
	"bluebird/usecase"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	dbType := "mysql"
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "bluebird"

	db, err := gorm.Open(dbType, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	router := mux.NewRouter().StrictSlash(true)

	// userRepository := repository.NewUser(db)
	// cartRepository := repository.NewCart(db)
	// orderRepository := repository.NewOrder(db)
	productRepository := repository.NewProduct(db)

	productUsecase := usecase.NewProduct(productRepository)
	api.NewProductHandler(router, productUsecase)

	log.Fatal(http.ListenAndServe(":10000", nil))

}
