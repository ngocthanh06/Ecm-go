package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"net/http"
)

func init() {
	utils.Route = "admin"
}

func AdminRoutes(router *gin.Engine) {
	routeAdmin := router.Group("/admin")
	{
		routeAdmin.GET("/", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "index", gin.H{"title": "Admin"}, true)
		})
	}

	return
}
