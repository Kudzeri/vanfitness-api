package main

import (
	"github.com/Kudzeri/vanfitness-api/config"
	"github.com/Kudzeri/vanfitness-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	r := gin.Default()

	r.Use(config.SetupCORS())

	api := r.Group("/api")
	routes.SetupAuthRoutes(api)
	routes.SetupProfileRoutes(api)
	routes.SetupUserRoutes(api)

	r.Run(":8080")
}
