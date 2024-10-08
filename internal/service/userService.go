package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ngocthanh06/ecommerce/internal/config"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
	"github.com/ngocthanh06/ecommerce/internal/repository"
	validation "github.com/ngocthanh06/ecommerce/internal/validation/user"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserServiceInterface interface {
	HomeList(ctx context.Context) ([]map[string]interface{}, error)
	Register(params *validation.RegisterData)
	VerifyUserInformation(token string) (*model.User, error)
}

type UserService struct {
	userRepo *repository.UserRepository
}

type homeResponseType struct {
	Categories []*model.Category
	Products   []*model.Product
}

// NewUserService
//
// Parameters:
//
// Returns:
// - *UserService
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.GetRepository().UserRepository,
	}
}

// HomeList
//
// Parameters:
// - ctx: *content.Context
//
// Returns:
// - *homeResponseType
// - error
func (userRepo UserService) HomeList(ctx context.Context) (*homeResponseType, error) {
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

// Register
//
// Parameters:
// - user: *model.User
// - token: string
//
// Returns:
// - *model.User
// - error
func (userRepo UserService) Register(user *model.User, token string) (*model.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("hash pass error")

		return nil, err
	}
	user.Password = string(hashPass)
	user.Role = utils.Roles["user"]
	userJSON, err := json.Marshal(user)

	if err != nil {
		return nil, err
	}

	// set value in redis 30 minutes
	err = database.GetRedisInstance().Set(database.CtxBg, "registration:"+token, userJSON, time.Minute*30).Err()

	if err != nil {
		fmt.Printf("registration fails: %v", err)
	}

	err = database.GetRedisInstance().Set(database.CtxBg, "registration:"+user.Email, user.Email, time.Minute*30).Err()

	if err != nil {
		fmt.Printf("registration fails: %v", err)
	}

	// send Email register
	sendEmailVerifyRegisterAccount(user.Email)

	return user, nil
}

// SendEmailVerifyUserRegisterAccount
//
// Parameters:
// - email: string
// - name: string
//
// Returns:
// - bool
func sendEmailVerifyRegisterAccount(email string) bool {
	emailList := config.SetEmailAddress([]string{email})

	file := string(utils.RenderHtmlTemplateMail("user/verify-register-account"))

	// test send mail
	sendMail := config.SetMailContent(
		emailList,
		"You are awesome!",
		file,
		nil,
		emailList,
		emailList,
	)

	return sendMail.SendMailWithMailtrap()
}

// VerifyUserInformation
//
// Parameters:
// - userRedis: string
//
// Returns:
// - *model.User
// - error
func (userRepo UserService) VerifyUserInformation(userRedis string) (*model.User, error) {
	var user *model.User
	err := json.Unmarshal([]byte(userRedis), &user)

	if err != nil {
		fmt.Println(err)

		return nil, nil
	}

	// handle logic
	err = userRepo.userRepo.FirstUser(user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// create
	err = userRepo.userRepo.CreateUser(user).Error

	//
	if err != nil {
		return nil, err
	}

	return user, nil
}
