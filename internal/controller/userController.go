package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/internal/database"
	"github.com/ngocthanh06/ecommerce/internal/service"
	validation "github.com/ngocthanh06/ecommerce/internal/validation/user"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"log"
	"net/http"
)

// Home
//
// Parameters:
// - context: *gin.Context
func Home(context *gin.Context) {
	response, err := service.GetServiceInstance().UserService.HomeList(context)

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	// response FE
	utils.TemplateRender(context, http.StatusOK, "indexUser", gin.H{
		"title":      "User",
		"categories": response.Categories,
	})

	return
}

// Register
//
// Parameters:
// - ctx: *gin.Context
func Register(ctx *gin.Context) {
	var validate validation.RegisterData
	user, errors := validate.ValidationRequest(ctx)

	if errors != nil {
		utils.TemplateRender(ctx, http.StatusBadRequest, "userRegister", gin.H{
			"title":  "User",
			"params": user,
			"errors": errors,
		})

		return
	}

	token, _ := utils.CreateRandToken(32)

	_, err := service.GetServiceInstance().UserService.Register(user, token)

	if err != nil {
		utils.TemplateRender(ctx, http.StatusBadRequest, "userRegister", gin.H{
			"title":  "User",
			"params": user,
			"errors": err,
		})

		return
	}

	ctx.Redirect(http.StatusFound, "/page-verify-account-register/"+token)

	return
}

// VerifyUserInformation
//
// Parameters:
// - ctx: *gin.Context
func VerifyUserInformation(ctx *gin.Context) {
	token := ctx.Param("token")

	// check token registration exist in redis
	userRedis, err := database.GetRedisInstance().Get(ctx, "registration:"+token).Result()
	if err != nil {
		fmt.Println("Registration not found")

		ctx.Redirect(http.StatusFound, "/404")
		return
	}

	user, err := service.GetServiceInstance().UserService.VerifyUserInformation(userRedis)

	if err != nil {
		fmt.Println("Registration not found")

		ctx.Redirect(http.StatusFound, "/404")
		return
	}

	// remove redis
	database.GetRedisInstance().Del(ctx, "registration:"+user.Email)
	database.GetRedisInstance().Del(ctx, "registration:"+token)

	ctx.Redirect(http.StatusFound, "/login")

	return
}

// ShowPageVerifyAccount
//
// Parameters:
// - ctx: *gin.Context
func ShowPageVerifyAccount(ctx *gin.Context) {
	token := ctx.Param("token")

	// check token registration exist in redis
	_, err := database.GetRedisInstance().Get(ctx, "registration:"+token).Result()

	if err != nil {
		ctx.Redirect(http.StatusFound, "/404")
		return
	}

	utils.TemplateRender(ctx, http.StatusOK, "pageVerifyAccountRegister", gin.H{
		"title": "Please check email verify account!",
	})
}
