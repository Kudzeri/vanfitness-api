package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/user", middleware.AuthMiddleware(), controllers.GetUser)
}