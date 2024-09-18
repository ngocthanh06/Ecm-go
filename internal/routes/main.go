package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ngocthanh06/ecommerce/internal/repository"
	"github.com/ngocthanh06/ecommerce/internal/service"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"log"
)

func init() {
	utils.ExtFile = "html"
}

// MainRoutes
//
// Parameters:
//
// Returns:
func MainRoutes() {
	router := gin.Default()

	templates, err := utils.LoadTemplateDir("web/templates/")

	if err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	router.SetHTMLTemplate(templates)
	router.Static("/assets", "./web/assets")

	// repository
	repository.InitBaseRepository()

	// service
	service.InitService()

	// user route
	UserRoutes(router)

	// admin route
	AdminRoutes(router)

	router.Run(fmt.Sprintf(":%s", "8001"))
}
