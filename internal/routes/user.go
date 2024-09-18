package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/internal/controller"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"net/http"
)

// UserRoutes
//
// Parameters:
// - route: *gin.Engine
//
// Returns:
func UserRoutes(route *gin.Engine) {
	// Page not found
	//
	// Parameters:
	// - context: *gin.Context
	route.GET("404", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userPageNotFound", gin.H{
			"title": "User page not found",
		})
	})

	// Register
	//
	// Parameters:
	// - context: *gin.Context
	route.GET("/register", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userRegister", gin.H{
			"title": "Register User",
		})
	})

	// Register post
	//
	// Parameters:
	// - context: *gin.Context
	route.POST("/register", func(context *gin.Context) {
		controller.Register(context)
	})

	// Page not found
	//
	// Parameters:
	// - context: *gin.Context
	route.GET("/verify/:token", func(context *gin.Context) {
		controller.VerifyUserInformation(context)
	})

	// Page verify account register :token
	//
	// Parameters:
	// - context: *gin.Context
	route.GET("/page-verify-account-register/:token", func(context *gin.Context) {
		controller.ShowPageVerifyAccount(context)
	})

	// Login
	//
	// Parameters:
	// - context: *gin.Context
	route.GET("/login", func(context *gin.Context) {
		utils.TemplateRender(context, http.StatusOK, "userLogin", gin.H{
			"title": "Login",
		})
	})

	// Group user
	//
	// Parameters:
	// - context: *gin.Context
	userRoute := route.Group("user")
	{
		// Register
		//
		// Parameters:
		// - context: *gin.Context
		userRoute.GET("/register", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userRegister", gin.H{
				"title": "Register User",
			})
		})

		// Login
		//
		// Parameters:
		// - context: *gin.Context
		userRoute.GET("/login", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userLogin", gin.H{
				"title": "Login",
			})
		})

		// get User
		//
		// Parameters:
		// - context: *gin.Context
		userRoute.GET("/", func(context *gin.Context) {
			// call controller and response data
			controller.Home(context)
		})

		// shop-cart
		//
		// Parameters:
		// - context: *gin.Context
		userRoute.GET("/shop-cart", func(context *gin.Context) {
			utils.TemplateRender(context, http.StatusOK, "userShopCart", gin.H{
				"title": "Shop Cart",
			})
		})
	}
}
