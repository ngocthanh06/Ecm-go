package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
	validation "github.com/ngocthanh06/ecommerce/internal/validation/user"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	HomeList(ctx context.Context) ([]map[string]interface{}, error)
	Register(params *validation.RegisterData)
}

type UserService struct {
	repo *database.Database
}

type homeResponseType struct {
	Categories []*model.Category
	Products   []*model.Product
}

func NewUserService() *UserService {
	return &UserService{
		repo: database.GetDb(),
	}
}

func (dbStorage UserService) HomeList(ctx context.Context) (*homeResponseType, error) {
	categories := []*model.Category{}
	// get categories
	resultCategories, err := GetCategories(categories)

	if err != nil {
		fmt.Printf("Error fetching categories: %v", err)

		return nil, err
	}

	response := &homeResponseType{
		Categories: resultCategories,
	}

	return response, nil
}

func (dbStorage UserService) Register(params *validation.RegisterData) (*model.User, error) {
	// handle logic
	err := dbStorage.repo.Db.Table("users").
		Where("email = ? AND phone = ?", params.Email, params.Phone).
		First(&params).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// send data to redis

	// create hash password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("hash pass error")

		return nil, err
	}

	params.Password = string(hashPass)

	var user = &model.User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  params.Password,
		Phone:     params.Phone,
		Role:      utils.Roles["user"],
	}
	// create
	result := dbStorage.repo.Db.Table("users").Create(user)

	if result.Error != nil {
		fmt.Println(user)

		return nil, result.Error
	}

	return user, nil
}
