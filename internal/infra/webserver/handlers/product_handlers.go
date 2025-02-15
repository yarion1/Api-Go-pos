package handlers

import (
	"encoding/json"
	"net/http"
	"pos-go-api/internal/dto"
	"pos-go-api/internal/entity"
	"pos-go-api/internal/infra/database"
)

type ProductHandler struct {
	productDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		productDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.productDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
