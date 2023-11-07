package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nathangds/order-api/models"
)

var categories []models.Category = []models.Category{
	*models.NewCategory(1, "description", models.CLOTHES),
}

func ShowCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)
}

func AddCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestData models.NewCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(requestData)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, e.Field()+" is "+e.Tag())
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessages)
		return
	}

	var newCategory = *models.NewCategory(int64(len(categories)+1), requestData.Description, requestData.CategoryType)
	categories = append(categories, newCategory)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCategory)
}
