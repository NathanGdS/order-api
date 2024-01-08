package models

import (
	"github.com/google/uuid"
)

type NewItemRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	CategoryId  string `json:"category_id" validate:"required"`
}

type Item struct {
	ItemId      string   `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"item_id"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"optional"`
	CategoryId  string   `json:"category_id" validate:"required"`
	Category    Category `gorm:"foreignKey:CategoryId;references:CategoryId" json:"category"`
	BaseModels
}

func NewItem(name string, description string, categoryId string) *Item {
	var id = uuid.New()
	return &Item{
		ItemId:      id.String(),
		Name:        name,
		Description: description,
		CategoryId:  categoryId,
	}
}
