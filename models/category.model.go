package models

import (
	"time"

	"github.com/google/uuid"
)

type CategoryTypes string

type NewCategoryRequest struct {
	Description  string        `json:"description" validate:"required"`
	CategoryType CategoryTypes `json:"category_type" validate:"required"`
}

type UpdateCategoryRequest struct {
	Description  string        `json:"description"`
	CategoryType CategoryTypes `json:"category_type"`
}

const (
	CLOTHES CategoryTypes = "clothes"
	FOOD    CategoryTypes = "food"
	TECH    CategoryTypes = "tech"
	BOOKS   CategoryTypes = "books"
)

type Category struct {
	CategoryId   string        `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"category_id"`
	Description  string        `json:"description" validate:"required"`
	CategoryType CategoryTypes `json:"category_type" validate:"required"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    time.Time     `json:"deleted_at"`
	Items        []Item        `json:"items" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewCategory(description string, categoryType CategoryTypes) *Category {
	var id = uuid.New()
	return &Category{
		CategoryId:   id.String(),
		Description:  description,
		CategoryType: categoryType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Time{},
	}
}
