package routers

import (
	"ims/internal/interfaces/http/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewRouter(dataBase *pgxpool.Pool) *mux.Router {

	productHandler := handlers.NewProductHandler(dataBase)   // Изменено
	categoryHandler := handlers.NewCategoryHandler(dataBase) // Изменено

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
