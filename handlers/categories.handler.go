package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nathangds/order-api/factories"
	"github.com/nathangds/order-api/helpers"
	"github.com/nathangds/order-api/models"
)

func (h Handler) ShowCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category

	if result := h.DB.Preload("Items").Find(&categories); result.Error != nil {
		factories.ResponseFactory(w, http.StatusInternalServerError, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, categories)
}

func (h Handler) ShowCategoryById(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	vars := mux.Vars(r)
	categoryId := vars["id"]

	if result := h.DB.First(&category, "category_id = ?", categoryId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	var filteredResults []models.Item = make([]models.Item, 0)

	for _, item := range category.Items {
		if item.DeletedAt.IsZero() {
			filteredResults = append(filteredResults, item)
		}
	}

	factories.ResponseFactory(w, http.StatusOK, filteredResults)
}

func (h Handler) AddCategories(w http.ResponseWriter, r *http.Request) {
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

	var newCategory = *models.NewCategory(requestData.Description, requestData.CategoryType)

	if result := h.DB.Create(&newCategory); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusCreated, newCategory)
}

func (h Handler) UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	vars := mux.Vars(r)
	categoryId := vars["id"]
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

	if result := h.DB.First(&category, "category_id = ?", categoryId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	if requestData.Description != "" {
		category.Description = requestData.Description
	}

	if requestData.CategoryType != "" {
		category.CategoryType = requestData.CategoryType
	}

	category.UpdatedAt = time.Now()

	if result := h.DB.Updates(&category).Where("category_id = ?", categoryId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, category)
}

func (h Handler) RemoveCategoryById(w http.ResponseWriter, r *http.Request) {
	var category *models.Category
	vars := mux.Vars(r)
	categoryId := vars["id"]

	if result := h.DB.First(&category, "category_id = ?", categoryId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	category.DeletedAt = time.Now()

	if result := h.DB.Save(&category); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusNoContent, nil)
}
