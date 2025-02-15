package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"pos-go-api/configs"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	"pos-go-api/internal/infra/webserver/handlers"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	ProductDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(ProductDB)

	http.HandleFunc("/products", ProductHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
