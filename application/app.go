package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}



func New() *App {
	app := &App{
	}
	app.loadRoutes()
	return app
}

func (a *App) Start(ctx context.Context, addr string) error  {
	server := &http.Server{
		Addr: ":" + addr,
		Handler: a.router,
	}	
	defer server.Close()
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server %w", err)
	}
	return nil
}