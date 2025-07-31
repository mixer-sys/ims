package handlers

import (
	"encoding/json"
	"ims/internal/domain/models"
	"ims/internal/domain/ports/repository"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	db repository.CategoryRepository
}

func NewCategoryHandler(db repository.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{db: db}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.ID = uuid.New().String()
	err := h.db.Create(r.Context(), &category)
	if err != nil {
		http.Error(w, "Error inserting category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	category, err := h.db.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(category); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.ID = id
	err := h.db.Update(r.Context(), &category)
	if err != nil {
		http.Error(w, "Error updating category", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.db.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting category", http.StatusInternalServerError)
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
		http.Error(w, "Error fetching categories", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(categories); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
