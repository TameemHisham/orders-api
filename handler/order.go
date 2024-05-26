package handler

import (
	"fmt"
	"net/http"
)







type Order struct {	
	FirstName string 
	LastName string
	Price float64
	Availability bool
}

func (o *Order)  Create (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Order")
}
func (o *Order)  List (w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all Orders")
}

func (o *Order)  GetByID (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order by ID")
}
func (o *Order)  UpdateByID (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by ID")
}
func (o *Order)  DeleteByID (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by ID")
}