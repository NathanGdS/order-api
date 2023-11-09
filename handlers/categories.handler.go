package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/factories"
	"github.com/nathangds/order-api/models"
)

var categories []models.Category = []models.Category{
	*models.NewCategory(1, "description", models.CLOTHES),
}

func ShowCategories(w http.ResponseWriter, r *http.Request) {
	factories.ResponseFactory(w, http.StatusOK, categories)
}

func ShowCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var category models.Category

	for _, c := range categories {
		if strconv.Itoa(int(c.CategoryId)) == vars["id"] {
			category = c
			break
		}
	}

	if category.CategoryId == 0 {
		factories.ResponseFactory(w, http.StatusNotFound, factories.ErrorResponse([]string{"Category not found"}))
		return
	}
	factories.ResponseFactory(w, http.StatusOK, category)
}

func AddCategories(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestData models.NewCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Invalid JSON request body"}))
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
		factories.ResponseFactory(w, http.StatusNotFound, factories.ErrorResponse(errorMessages))
		return
	}

	var newCategory = *models.NewCategory(int64(len(categories)+1), requestData.Description, requestData.CategoryType)
	categories = append(categories, newCategory)
	factories.ResponseFactory(w, http.StatusCreated, newCategory)
}
