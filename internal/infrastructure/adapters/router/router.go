package routers

import (
	"ims/internal/interfaces/http/handlers"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewRouter(dataBase *pgxpool.Pool) *mux.Router {
	productHandler := handlers.NewProductHandler(dataBase)
	categoryHandler := handlers.NewCategoryHandler(dataBase)

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
