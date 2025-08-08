package handlers

import (
	"context"
	"encoding/json"
	"ims/internal/domain/errors"
	"ims/internal/domain/models"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, limit, offset int) ([]models.Product, error)
}

type ProductHandler struct {
	db ProductRepository
}

func NewProductHandler(db ProductRepository) *ProductHandler {
	return &ProductHandler{db: db}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		apiErr := errors.NewAPIError("Invalid request payload", http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	err := models.ValidateProduct(&product)
	if err != nil {
		apiErr := errors.NewAPIError("Invalid product data", http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	if err := h.db.Create(r.Context(), &product); err != nil {
		apiErr := errors.NewAPIError("Error inserting product", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := h.db.GetByID(r.Context(), id)
	if err != nil {
		apiErr := errors.NewAPIError("Product not found", http.StatusNotFound)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	if err := json.NewEncoder(w).Encode(product); err != nil {
		apiErr := errors.NewAPIError("Error encoding response", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		apiErr := errors.NewAPIError("Invalid request payload", http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	product.ID = id
	if err := h.db.Update(r.Context(), &product); err != nil {
		apiErr := errors.NewAPIError("Error updating product", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.db.Delete(r.Context(), id); err != nil {
		apiErr := errors.NewAPIError("Error deleting product", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	products, err := h.db.GetAll(r.Context(), limit, offset)
	if err != nil {
		apiErr := errors.NewAPIError("Error fetching products", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		apiErr := errors.NewAPIError("Error encoding response", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}
}
