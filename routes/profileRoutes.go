package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(r *gin.RouterGroup) {
	auth := r.Group("/profile")

	auth.GET("/create", middleware.AuthMiddleware(), controllers.MakeProfile)
}
