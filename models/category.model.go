package models

import "time"

type CategoryTypes string

const (
	CLOTHES CategoryTypes = "clothes"
	FOOD    CategoryTypes = "food"
	TECH    CategoryTypes = "tech"
	BOOKS   CategoryTypes = "books"
)

type Category struct {
	CategoryId   string
	Description  string
	CategoryType CategoryTypes
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

func NewCategory(categoryId string, description string, categoryType CategoryTypes) *Category {
	return &Category{
		CategoryId:   categoryId,
		Description:  description,
		CategoryType: categoryType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    time.Time{},
	}
}

type CategoryResponseDto struct {
	CategoryId   string        `json:"category_id"`
	Description  string        `json:"description"`
	CategoryType CategoryTypes `json:"category_type"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    interface{}   `json:"deleted_at"`
}

func CreateCategoryResponseDto(category *Category) CategoryResponseDto {
	var deletedAt = checkDate(category.DeletedAt)

	return CategoryResponseDto{
		CategoryId:   category.CategoryId,
		Description:  category.Description,
		CategoryType: category.CategoryType,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

func checkDate(date time.Time) interface{} {
	if date.IsZero() {
		return nil
	} else {
		return date.Format("2006-01-02 15:04:05")
	}
}
