package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"ims/internal/domain/errors"
	"ims/internal/domain/models"

	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *models.Category) error
	GetByID(ctx context.Context, id string) (*models.Category, error)
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context, limit, offset int) ([]models.Category, error)
}

type CategoryHandler struct {
	db CategoryRepository
}

func NewCategoryHandler(db CategoryRepository) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		apiErr := errors.NewAPIError("Invalid request payload", http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	err := models.ValidateCategory(&category)
	if err != nil {
		apiErr := errors.NewAPIError(fmt.Sprintf("Invalid category: %s", err), http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	category.ID = uuid.New().String()
	err = h.db.Create(r.Context(), &category)

	if err != nil {
		apiErr := errors.NewAPIError("Error inserting category", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	category, err := h.db.GetByID(r.Context(), id)
	if err != nil {
		apiErr := errors.NewAPIError("Category not found", http.StatusNotFound)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	if err := json.NewEncoder(w).Encode(category); err != nil {
		apiErr := errors.NewAPIError("Error encoding response", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		apiErr := errors.NewAPIError("Invalid request payload", http.StatusBadRequest)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	category.ID = id
	err := h.db.Update(r.Context(), &category)

	if err != nil {
		apiErr := errors.NewAPIError("Error updating category", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.db.Delete(r.Context(), id)
	if err != nil {
		apiErr := errors.NewAPIError("Error deleting category", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
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

	categories, err := h.db.GetAll(r.Context(), limit, offset)
	if err != nil {
		apiErr := errors.NewAPIError("Error fetching categories", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		apiErr := errors.NewAPIError("Error encoding response", http.StatusInternalServerError)
		errors.WriteErrorResponse(w, apiErr)
		return
	}
}
