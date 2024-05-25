package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
var counter int = 0

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)
	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}	
	defer server.Close()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while Listen to Serve, Error: %v", err)
	}
}



func basicHandler (w http.ResponseWriter, r *http.Request) {
	counter++
	w.Write([]byte(fmt.Sprintf("Hi Tameem! %d", counter)))
}