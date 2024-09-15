package validation

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ngocthanh06/ecommerce/internal/validation"
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

func (params *RegisterData) ValidationRequest(ctx *gin.Context) (*RegisterData, map[string]string) {
	if err := ctx.ShouldBind(&params); err != nil {
		fmt.Println(err)

		return nil, params.MessagesError(err)
	}

	if params.Password != params.ConfirmPassword {
		return nil, map[string]string{
			"ConfirmPassword": "Password not matching!",
		}
	}

	return params, nil
}

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
