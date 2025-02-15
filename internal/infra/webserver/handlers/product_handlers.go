package handlers

import (
	"github.com/gofiber/fiber/v2"
	"pos-go-api/internal/dto"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
	entityPkg "pos-go-api/pkg/entity"
	"strconv"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDB: db,
	}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product dto.CreateProductInput

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product data"})
	}

	if err := h.productDB.Create(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save product"})
	}

	return c.Status(fiber.StatusCreated).JSON(p)
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 10
	}
	sort := c.Query("sort")

	products, err := h.productDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get products"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": products})

}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	product, err := h.productDB.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product not found"})
	}
	c.GetRespHeader("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	var product entity.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id not found"})
	}

	_, err = h.productDB.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	err = h.productDB.Update(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}
	_, err := h.productDB.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	err = h.productDB.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Server error"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
}
