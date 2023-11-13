package models

import "time"

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
	CategoryId   int64         `json:"category_id"`
	Description  string        `json:"description" validate:"required"`
	CategoryType CategoryTypes `json:"category_type" validate:"required"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    time.Time     `json:"deleted_at"`
}

func NewCategory(categoryId int64, description string, categoryType CategoryTypes) *Category {
	return &Category{
		CategoryId:   categoryId,
		Description:  description,
		CategoryType: categoryType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Time{},
	}
}
