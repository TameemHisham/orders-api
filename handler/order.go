package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)
	

var rend = render.New() // pass options if you want


type Shop struct {
	DB *sql.DB
}

type Item struct {
    Name         string
    Price        float64
    Availability bool
}



func (s *Shop) Create(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    priceStr := chi.URLParam(r, "price")
    availabilityStr := chi.URLParam(r, "availability")

    // Convert string values to their respective types
    price, err := strconv.ParseFloat(priceStr, 64)
    if err != nil {
        http.Error(w, "Invalid price", http.StatusBadRequest)
        return
    }

    availability, err := strconv.ParseBool(availabilityStr)
    if err != nil {
        http.Error(w, "Invalid availability", http.StatusBadRequest)
        return
    }

    // Create the Item instance
    itemCreate := &Item{
        Name:         name,
        Price:        price,
        Availability: availability,
    }

    // Insert the Item into the database
    query := `INSERT INTO product (name, price, availability) VALUES ($1, $2, $3) RETURNING id`
		var pk int;
    err = s.DB.QueryRow(query, itemCreate.Name, itemCreate.Price, itemCreate.Availability).Scan(&pk)
    if err != nil {
        http.Error(w, "Error inserting item into database", http.StatusInternalServerError)
        return
    }

    // Now you can use the itemCreate variable as needed
    fmt.Printf("Item created: %+v\n", *itemCreate)
		rend.JSON(w, 200, itemCreate)

}

func (s *Shop) List(w http.ResponseWriter, r *http.Request)  {
	data := []Item{}
	rows, err := s.DB.Query("SELECT name, availability, price FROM product")
	if err != nil {
		http.Error(w, "Error finding items in database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Println("List all Items")
	var Name         string
	var Price        float64
	var Availability bool
	for rows.Next() {
		err := rows.Scan(&Name, &Availability, &Price)
		if err != nil {
			http.Error(w, "Error iterating to find items in database", http.StatusInternalServerError)
			return
		}
		data = append(data, Item{Name, Price, Availability})
	}
	fmt.Println("List All Items")
	rend.JSON(w, 200, data)
	
}

func (s *Shop) GetByID(w http.ResponseWriter, r *http.Request){
	ID := chi.URLParam(r, "id")
	wantedItem := &Item{}
	query := `SELECT name, price, availability FROM product WHERE id = $1`
	err := s.DB.QueryRow(query, ID).Scan(&wantedItem.Name, &wantedItem.Price, &wantedItem.Availability)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error finding item by id in database (ID %v may not EXISTS)", ID), http.StatusInternalServerError)
	} 			
	fmt.Println("Get an Item by ID")
	rend.JSON(w, 200, wantedItem)
}

func (s *Shop) UpdateByID(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	priceStr := chi.URLParam(r, "price")
	availabilityStr := chi.URLParam(r, "availability")
	id := chi.URLParam(r, "id")
	// Convert string values to their respective types
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
	}

	availability, err := strconv.ParseBool(availabilityStr)
	if err != nil {
			http.Error(w, "Invalid availability", http.StatusBadRequest)
			return
	}

	// Create the Item instance
	itemCreate := &Item{
			Name:         name,
			Price:        price,
			Availability: availability,
	}

	// Insert the Item into the database
	query := `UPDATE product SET name=$1, price=$2, availability=$3 WHERE id=$4 RETURNING name, price, availability`

	newData := &Item{}
	err = s.DB.QueryRow(query, itemCreate.Name, itemCreate.Price, itemCreate.Availability, id).Scan(&newData.Name, &newData.Price, &newData.Availability)

	if err != nil {
			http.Error(w, "Error inserting item into database", http.StatusInternalServerError)
			fmt.Print("Error: %w", err)
			return
	}
	// Now you can use the itemCreate variable as needed
	fmt.Printf("Item Updated: %+v\n", *itemCreate)
	rend.JSON(w, 200, newData)


	
	
	fmt.Println("Update an Item by ID")
}

func (s *Shop) DeleteByID(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		query := `DELETE FROM product WHERE id=$1;`
	
		// Exec returns a result and an error, handle them correctly
		result, err := s.DB.Exec(query, id)
		if err != nil {
			http.Error(w, "Error deleting item from database", http.StatusInternalServerError)
			fmt.Printf("Error deleting item: %v\n", err)
			return
		}
	
		// Optionally, check how many rows were affected
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(w, "Error retrieving affected rows", http.StatusInternalServerError)
			fmt.Printf("Error retrieving affected rows: %v\n", err)
			return
		}
	
		// Inform the user if no rows were affected (e.g., item not found)
		if rowsAffected == 0 {
			http.Error(w, "No item found with the provided ID", http.StatusNotFound)
			return
		}
	
		// Respond with success if the deletion was successful
		rend.JSON(w, 200, result)

		fmt.Println("Delete an Item by ID")
	
}
