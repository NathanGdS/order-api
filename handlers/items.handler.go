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

func (h Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var requestData models.NewItemRequest

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

	if categoryExists := h.DB.First(&models.Category{}, "category_id = ?", requestData.CategoryId); categoryExists.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Category not found"}))
		return
	}

	var newItem = *models.NewItem(requestData.Name, requestData.Description, requestData.CategoryId, requestData.Value)

	if result := h.DB.Create(&newItem); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusCreated, newItem.ShowItem())

}

func (h Handler) ShowItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item

	if result := h.DB.Find(&items); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	var filteredResults []models.ResponseItem = make([]models.ResponseItem, 0)

	for _, item := range items {
		if item.DeletedAt.IsZero() {
			filteredResults = append(filteredResults, item.ShowItem())
		}
	}

	factories.ResponseFactory(w, http.StatusOK, filteredResults)
}

func (h Handler) ShowById(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	vars := mux.Vars(r)
	itemId := vars["id"]

	if result := h.DB.First(&item, "item_id = ?", itemId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Item not found"}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, item.ShowItem())

}

func (h Handler) UpdateItemById(w http.ResponseWriter, r *http.Request) {
	var item *models.Item
	vars := mux.Vars(r)
	itemId := vars["id"]
	defer r.Body.Close()

	var requestData models.UpdateItemRequest

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Invalid JSON request body"}))
		return
	}

	if requestData == (models.UpdateItemRequest{}) {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"You cannot update an item with an empty request body"}))
		return
	}

	if result := h.DB.First(&item, "item_id = ?", itemId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusNotFound, factories.ErrorResponse([]string{"Item not found"}))
		return
	}

	if requestData.Name != "" {
		item.Name = requestData.Name
	}

	if requestData.Description != "" {
		item.Description = requestData.Description
	}

	if requestData.Value != 0 {
		item.Value = helpers.DecimalToCents(requestData.Value)
	}

	var category models.Category

	if requestData.CategoryId != "" {
		if categoryExists := h.DB.First(&category, "category_id = ?", requestData.CategoryId); categoryExists.Error != nil || !category.DeletedAt.IsZero() {
			factories.ResponseFactory(w, http.StatusNotFound, factories.ErrorResponse([]string{"Category not found"}))
			return
		}
		item.CategoryId = requestData.CategoryId
	}

	if result := h.DB.Updates(&item).Where("item_id = ?", itemId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, item.ShowItem())
}

func (h Handler) RemoveItemById(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	vars := mux.Vars(r)
	itemId := vars["id"]

	if result := h.DB.First(&item, "item_id = ?", itemId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Item not found"}))
		return
	}

	item.DeletedAt = time.Now()

	if result := h.DB.Save(&item); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, item.ShowItem())
}
