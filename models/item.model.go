package models

import (
	"github.com/google/uuid"
	"github.com/nathangds/order-api/helpers"
)

type NewItemRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	CategoryId  string  `json:"category_id" validate:"required"`
	Value       float64 `json:"value" validate:"required"`
}

type UpdateItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  string  `json:"category_id"`
	Value       float64 `json:"value"`
}

type ResponseItem struct {
	ItemId      string  `json:"item_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryId  string  `json:"category_id"`
	Value       float64 `json:"value"`
	BaseModels
}

type Item struct {
	ItemId      string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"item_id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"optional"`
	CategoryId  string `json:"category_id" validate:"required"`
	Value       int64  `json:"value" validate:"required"`
	BaseModels
}

func NewItem(name string, description string, categoryId string, value float64) *Item {
	var id = uuid.New()
	return &Item{
		ItemId:      id.String(),
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
		Value:       helpers.DecimalToCents(value),
	}
}

func (i *Item) ShowItem() ResponseItem {
	return ResponseItem{
		ItemId:      i.ItemId,
		Name:        i.Name,
		Description: i.Description,
		CategoryId:  i.CategoryId,
		Value:       helpers.CentsToDecimal(i.Value),
		BaseModels:  i.BaseModels,
	}
}
