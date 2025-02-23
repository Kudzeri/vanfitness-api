package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup) {
	auth := r.Group("/user")

	auth.GET("/me", middleware.AuthMiddleware(), controllers.GetUser)
	auth.PUT("/update", middleware.AuthMiddleware(), controllers.UpdateUser)
}
