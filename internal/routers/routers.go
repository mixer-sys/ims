package routers

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"ims/internal/db"
	"ims/internal/handlers"
	services "ims/internal/services"
)

func NewRouter(dataBase *sql.DB) *mux.Router {
	productRepo := db.NewProductRepository(dataBase)
	categoryRepo := db.NewCategoryRepository(dataBase)

	productService := services.NewProductService(productRepo)
	categoryService := services.NewCategoryService(categoryRepo)

	productHandler := handlers.NewProductHandler(productService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	r := mux.NewRouter()
	r.HandleFunc("/products", productHandler.Create).Methods("POST")
	r.HandleFunc("/products/{id:[0-9]+}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/products/{id:[0-9]+}", productHandler.Update).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9]+}", productHandler.Delete).Methods("DELETE")
	r.HandleFunc("/products", productHandler.GetAll).Methods("GET")

	r.HandleFunc("/categories", categoryHandler.Create).Methods("POST")
	r.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.GetByID).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.Update).Methods("PUT")
	r.HandleFunc("/categories/{id:[0-9]+}", categoryHandler.Delete).Methods("DELETE")
	r.HandleFunc("/categories", categoryHandler.GetAll).Methods("GET")

	return r
}
