package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}	
	defer server.Close()
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error while Listen to Serve, Error: %v", err)
	}
}


func DisplayData(r *http.Request) {
	fmt.Printf("METHOD: %v\n", r.Method)
	fmt.Printf("REQUEST: %v\n", r.URL)
}

func basicHandler (w http.ResponseWriter, r *http.Request) {
	DisplayData(r)
	w.Write([]byte("Hi Tameem!"))
}