package service

import (
	"fmt"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
)

type CategoryService struct {
	repo *database.Database
}

// NewCategoryService
//
// Parameters:
//
// Returns:
// - *CategoryService
func NewCategoryService() *CategoryService {
	return &CategoryService{
		repo: database.GetDb(),
	}
}

// GetCategories
//
// Parameters:
// - categories: []*model.Category
//
// Returns:
// - []*model.Category
// - error
func GetCategories(categories []*model.Category) ([]*model.Category, error) {
	result := database.GetDb().Db.Find(&categories)

	if result.Error != nil {
		fmt.Printf("Error fetching categories: %v", result.Error)

		return nil, result.Error
	}

	return categories, nil
}
