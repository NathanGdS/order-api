package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nathangds/order-api/db"
	"github.com/nathangds/order-api/handlers"
	"github.com/nathangds/order-api/helpers"
	"github.com/nathangds/order-api/routes"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := ":" + helpers.GetEnvVariable("PORT")
	DB := db.Init()
	h := handlers.New(DB)
	router := &routes.Router{}
	router.Init().LoadHandlers(h)

	fmt.Println("Server running on port", PORT)
	http.ListenAndServe(PORT, router.Ref)
}
