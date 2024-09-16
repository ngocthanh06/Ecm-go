package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/internal/controller"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"net/http"
)

func init() {
	utils.Route = "user"
}

func UserRoutes(route *gin.Engine) {
	// Page not found
	route.GET("404", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userPageNotFound", gin.H{
			"title": "User page not found",
		})
	})

	// register.html
	route.GET("/register", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userRegister", gin.H{
			"title": "Register User",
		})
	})

	route.POST("/register", func(context *gin.Context) {
		controller.Register(context)
	})

	route.GET("/verify/:token", func(context *gin.Context) {
		controller.VerifyUserInformation(context)
	})

	route.GET("/page-verify-account-register/:token", func(context *gin.Context) {
		controller.ShowPageVerifyAccount(context)
	})

	// login.html
	route.GET("/login", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userLogin", gin.H{
			"title": "Login",
		})
	})

	userRoute := route.Group("user")
	{
		// register.html
		userRoute.GET("/register", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userRegister", gin.H{
				"title": "Register User",
			})
		})

		// login.html
		userRoute.GET("/login", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userLogin", gin.H{
				"title": "Login",
			})
		})

		userRoute.GET("/", func(context *gin.Context) {
			// call controller and response data
			controller.Home(context)
		})

		userRoute.GET("/shop-cart", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userShopCart", gin.H{
				"title": "Shop Cart",
			})
		})
	}
}
