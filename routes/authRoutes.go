package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.GET("/user", middleware.AuthMiddleware(), controllers.GetUser)
}
