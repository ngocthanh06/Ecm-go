package main

import (
	"github.com/joho/godotenv"
	"github.com/ngocthanh06/ecommerce/cmd"
	"github.com/ngocthanh06/ecommerce/internal/database"
	"github.com/ngocthanh06/ecommerce/internal/routes"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"log"
)

func init() {
	if err := godotenv.Load(utils.Env); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// connection database
	database.ConnectionDatabase()

	// init redis
	database.InitRedisClient()
}

func main() {
	// Execute cmd
	cmd.Execute()

	// connection route
	routes.MainRoutes()
}
