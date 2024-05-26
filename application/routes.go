package application

import (
	"database/sql"
	"net/http"

	"github.com/TameemHisham/orders-api/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes(router *chi.Mux, db *sql.DB) {
    router.Use(middleware.Logger)
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })
    router.Route("/orders", func(router chi.Router) {
        loadOrderRoutes(router, db)
    })
}

func loadOrderRoutes(router chi.Router, db *sql.DB) {
	orderHandler := &handler.Shop{DB: db} // Pass the db instance to the handler
	router.Post("/{name}-{price}-{availability}", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}/{name}-{price}-{availability}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
