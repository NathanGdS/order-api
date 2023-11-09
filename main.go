package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/handlers"
)

func main() {
	const PORT = ":4000"
	router := mux.NewRouter()

	router.HandleFunc("/categories", handlers.ShowCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", handlers.ShowCategoryById).Methods("GET")
	router.HandleFunc("/categories", handlers.AddCategories).Methods("POST")
	http.ListenAndServe(PORT, router)
	fmt.Println("Server running on port", PORT)
}
