package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/factories"
	"github.com/nathangds/order-api/helpers"
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

	val := helpers.ValidateRequest(requestData)
	if val != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse(val))
		return
	}

	var newCategory = *models.NewCategory(int64(len(categories)+1), requestData.Description, requestData.CategoryType)
	categories = append(categories, newCategory)
	factories.ResponseFactory(w, http.StatusCreated, newCategory)
}

func UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var category *models.Category

	for index, c := range categories {
		if strconv.Itoa(int(c.CategoryId)) == vars["id"] {
			category = &categories[index]
			break
		}
	}

	if category.CategoryId == 0 {
		factories.ResponseFactory(w, http.StatusNotFound, factories.ErrorResponse([]string{"Category not found"}))
		return
	}

	defer r.Body.Close()

	var requestData models.UpdateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Invalid JSON request body"}))
		return
	}

	if requestData == (models.UpdateCategoryRequest{}) {
		factories.ResponseFactory(w, http.StatusOK, category)
		return
	}

	if requestData.Description != "" {
		category.Description = requestData.Description
	}

	if requestData.CategoryType != "" {
		category.CategoryType = requestData.CategoryType
	}

	category.UpdatedAt = time.Now()

	factories.ResponseFactory(w, http.StatusOK, category)
}
