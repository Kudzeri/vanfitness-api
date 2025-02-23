package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupProfileRoutes(r *gin.RouterGroup) {
	profile := r.Group("/profile")

	profile.POST("/create", middleware.AuthMiddleware(), controllers.MakeProfile)
	profile.GET("/get", middleware.AuthMiddleware(), controllers.GetProfile)
	profile.PUT("/update",middleware.AuthMiddleware(),  controllers.UpdateProfile)

}
