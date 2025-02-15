package main

import (
	"github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"pos-go-api/configs"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	"pos-go-api/internal/infra/webserver/handlers"
)

func main() {
	config, err := configs.LoadConfig(".")
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

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.JWTSecret, config.JWTExpiresIn)

	f := fiber.New()
	f.Use(logger.New())
	f.Use(recover.New())

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(config.JWTSecret),
		},
		ContextKey: "user",
	})

	f.Route("/products", func(router fiber.Router) {
		router.Use(jwtMiddleware)
		router.Post("/", ProductHandler.CreateProduct)
		router.Get("/:id", ProductHandler.GetProduct)
		router.Get("/", ProductHandler.GetProducts)
		router.Put("/:id", ProductHandler.UpdateProduct)
		router.Delete("/:id", ProductHandler.DeleteProduct)
	})

	f.Post("/users", userHandler.Create)
	f.Post("/users/generate_token", userHandler.GetJWT)

	err = f.Listen(":8000")
	if err != nil {
		return
	}
}
