package handlers

import (
	"encoding/json"
	"net/http"

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

	var newItem = *models.NewItem(requestData.Name, requestData.Description, requestData.CategoryId)

	if result := h.DB.Create(&newItem); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusCreated, newItem)

}

func (h Handler) ShowItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item

	if result := h.DB.Preload("Category").Find(&items); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{result.Error.Error()}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, items)
}

func (h Handler) ShowById(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	vars := mux.Vars(r)
	itemId := vars["id"]

	if result := h.DB.Preload("Category").First(&item, "item_id = ?", itemId); result.Error != nil {
		factories.ResponseFactory(w, http.StatusBadRequest, factories.ErrorResponse([]string{"Item not found"}))
		return
	}

	factories.ResponseFactory(w, http.StatusOK, item)

}
