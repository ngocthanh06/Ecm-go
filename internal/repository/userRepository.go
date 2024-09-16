package repository

import (
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	repo *database.Database
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		repo: database.GetDb(),
	}
}

type UserRepositoryInterface interface {
	FirstUser(user *model.User) *gorm.DB
	CreateUser(user *model.User) *gorm.DB
}

func (dbStorage UserRepository) FirstUser(user *model.User) *gorm.DB {
	return dbStorage.repo.Db.Table("users").
		Where("email = ? OR phone = ?", user.Email, user.Phone).First(&user)
}

func (dbStorage UserRepository) CreateUser(user *model.User) *gorm.DB {
	return dbStorage.repo.Db.Table("users").Create(&user)
}
