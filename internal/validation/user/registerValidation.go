package validation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
	"github.com/ngocthanh06/ecommerce/internal/repository"
	"github.com/ngocthanh06/ecommerce/internal/validation"
	"gorm.io/gorm"
	"log"
)

type RegisterData struct {
	FirstName       string `form:"first_name" binding:"required"`
	LastName        string `form:"last_name" binding:"required"`
	Email           string `form:"email" binding:"required,email"`
	Phone           string `form:"phone" binding:"required"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required"`
}

type RegisterRequestInterface interface {
	ValidationRequest(ctx *gin.Context) (*RegisterData, map[string]string)
}

// ValidationRequest
//
// Parameters:
// - ctx: *gin.Context
//
// Returns:
// - *model.User
// - map[string]string
func (params *RegisterData) ValidationRequest(ctx *gin.Context) (*model.User, map[string]string) {
	if err := ctx.ShouldBind(params); err != nil {
		log.Println(err)

		return &model.User{}, params.MessagesError(err)
	}

	var user = &model.User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  params.Password,
		Phone:     params.Phone,
	}

	// check user exist in IB
	existUser := repository.GetRepository().UserRepository.FirstUser(user)

	if existUser.Error == nil || (existUser.Error != nil && existUser.Error != gorm.ErrRecordNotFound) {
		return user, map[string]string{
			"Account Exists": "Email or phone is exist!",
		}
	}

	// check account exist redis
	result, _ := database.GetRedisInstance().Get(ctx, "registration:"+params.Email).Result()
	if result != "" {
		return user, map[string]string{
			"AccountWaitVerify": "Email waiting verify please check email and accept link verify!",
		}
	}

	if params.Password != params.ConfirmPassword {
		return user, map[string]string{
			"ConfirmPassword": "Password not matching!",
		}
	}

	return user, nil
}

// MessagesError
//
// Parameters:
// - err: error
//
// Returns:
// - map[string]string
func (params *RegisterData) MessagesError(err error) map[string]string {
	errorsMessage := make(map[string]string)

	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorsMessage[fieldError.Field()] = validation.GetErrorMessage(fieldError.Tag(), fieldError.Field())
		}
	}

	return errorsMessage
}
