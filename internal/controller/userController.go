package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/internal/service"
	validation "github.com/ngocthanh06/ecommerce/internal/validation/user"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"log"
	"net/http"
)

func Home(context *gin.Context) {
	response, err := service.ServiceInstance.UserService.HomeList(context)

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	// response FE
	utils.TemplateRender(context, http.StatusOK, "indexUser", gin.H{
		"title":      "User",
		"categories": response.Categories,
	})
}

func Register(ctx *gin.Context) {
	var validate validation.RegisterData
	params, errors := validate.ValidationRequest(ctx)

	if errors != nil {
		utils.TemplateRender(ctx, http.StatusBadRequest, "userRegister", gin.H{
			"title":  "User",
			"errors": errors,
		})

		return
	}

	_, err := service.ServiceInstance.UserService.Register(params)

	if err != nil {
		utils.TemplateRender(ctx, http.StatusBadRequest, "userRegister", gin.H{
			"title":  "User",
			"errors": err,
		})

		return
	}

	utils.TemplateRender(ctx, http.StatusOK, "userRegister", gin.H{
		"title": "User",
	})
}
