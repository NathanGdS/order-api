package routes

import (
	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/handlers"
)

type Router struct {
	Ref *mux.Router
}

func (r *Router) Init() *Router {
	r.Ref = mux.NewRouter()
	return r
}

func (r *Router) LoadHandlers(h handlers.Handler) *mux.Router {

	r.Ref.HandleFunc("/categories", h.ShowCategories).Methods("GET")
	r.Ref.HandleFunc("/categories/{id}", h.ShowCategoryById).Methods("GET")
	r.Ref.HandleFunc("/categories", h.AddCategories).Methods("POST")
	r.Ref.HandleFunc("/categories/{id}", h.UpdateCategoryById).Methods("PUT")
	r.Ref.HandleFunc("/categories/{id}", h.RemoveCategoryById).Methods("DELETE")

	// Items
	r.Ref.HandleFunc("/item", h.CreateItem).Methods("POST")

	return r.Ref
}
