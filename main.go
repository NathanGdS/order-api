package main

import (
	"fmt"
	"net/http"

	"github.com/nathangds/order-api/db"
	"github.com/nathangds/order-api/handlers"
	"github.com/nathangds/order-api/helpers"
	"github.com/nathangds/order-api/routes"
)

func main() {
	PORT := ":" + helpers.GetEnvVariable("PORT")
	DB := db.Init()
	h := handlers.New(DB)
	router := &routes.Router{}
	router.Init().LoadHandlers(h)

	fmt.Println("Server running on port", PORT)
	http.ListenAndServe(PORT, router.Ref)
}
