package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/db"
	"github.com/nathangds/order-api/handlers"
	"github.com/nathangds/order-api/helpers"
)

func main() {
	PORT := ":" + helpers.GetEnvVariable("PORT")
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	// Categories
	router.HandleFunc("/categories", h.ShowCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", h.ShowCategoryById).Methods("GET")
	router.HandleFunc("/categories", h.AddCategories).Methods("POST")
	router.HandleFunc("/categories/{id}", h.UpdateCategoryById).Methods("PUT")
	router.HandleFunc("/categories/{id}", h.RemoveCategoryById).Methods("DELETE")

	// Items
	router.HandleFunc("/item", h.CreateItem).Methods("POST")

	fmt.Println("Server running on port", PORT)
	http.ListenAndServe(PORT, router)
}
