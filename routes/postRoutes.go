package routes

import (
	"github.com/Kudzeri/vanfitness-api/controllers"
	"github.com/Kudzeri/vanfitness-api/middleware"
	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.RouterGroup) {
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", middleware.AuthMiddleware(), controllers.CreatePost)
	r.PUT("/posts/:id", middleware.AuthMiddleware(), controllers.UpdatePost)
	r.DELETE("/posts/:id", middleware.AuthMiddleware(), controllers.DeletePost)
}
