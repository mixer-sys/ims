package routers

import (
	"ims/internal/domain/handlers"
	"ims/internal/domain/repository"

	"github.com/gorilla/mux"
)

func NewRouter(dataBase repository.Database) *mux.Router {
	productRepository := repository.NewProductRepository(&dataBase)
	categoryRepository := repository.NewCategoryRepository(&dataBase)

	productHandler := handlers.NewProductHandler(productRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	r := mux.NewRouter()

	r.HandleFunc("/products", productHandler.Create).Methods("POST")
	r.HandleFunc("/products/{id:[0-9a-fA-F-]+}", productHandler.GetByID).Methods("GET")
	r.HandleFunc("/products/{id:[0-9a-fA-F-]+}", productHandler.Update).Methods("PUT")
	r.HandleFunc("/products/{id:[0-9a-fA-F-]+}", productHandler.Delete).Methods("DELETE")
	r.HandleFunc("/products", productHandler.GetAll).Methods("GET")

	r.HandleFunc("/categories", categoryHandler.Create).Methods("POST")
	r.HandleFunc("/categories/{id:[0-9a-fA-F-]+}", categoryHandler.GetByID).Methods("GET")
	r.HandleFunc("/categories/{id:[0-9a-fA-F-]+}", categoryHandler.Update).Methods("PUT")
	r.HandleFunc("/categories/{id:[0-9a-fA-F-]+}", categoryHandler.Delete).Methods("DELETE")
	r.HandleFunc("/categories", categoryHandler.GetAll).Methods("GET")

	return r
}
