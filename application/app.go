package application

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	router http.Handler
}

func New(db *sql.DB) *App {
	app := &App{}
	app.setupRoutes(db)
	return app
}

func (a *App) setupRoutes(db *sql.DB) {
	router := chi.NewRouter()
	LoadRoutes(router, db)
	a.router = router
}

func (a *App) Start(ctx context.Context, addr string) error {
	server := &http.Server{
		Addr:    ":" + addr,
		Handler: a.router,
	}
	defer server.Close()
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server %w", err)
	}
	return nil
}
