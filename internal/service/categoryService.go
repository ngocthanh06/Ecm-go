package service

import (
	"fmt"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
)

type CategoryService struct {
	repo *database.Database
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		repo: database.GetDb(),
	}
}

func GetCategories(categories []*model.Category) ([]*model.Category, error) {
	result := database.GetDb().Db.Find(&categories)

	if result.Error != nil {
		fmt.Printf("Error fetching categories: %v", result.Error)

		return nil, result.Error
	}

	return categories, nil
}
