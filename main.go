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
	routes.SetupRoutes(r)
	r.Run(":8080")
}