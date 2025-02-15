package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	f := fiber.New()
	f.Use(logger.New())

	f.Post("/products", ProductHandler.CreateProduct)
	f.Get("/products/:id", ProductHandler.GetProduct)
	f.Get("/products", ProductHandler.GetProducts)
	f.Put("/products/:id", ProductHandler.UpdateProduct)
	f.Delete("/products/:id", ProductHandler.DeleteProduct)

	err = f.Listen(":8000")
	if err != nil {
		return
	}
}
