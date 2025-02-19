package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Kudzeri/vanfitness-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := utils.ValidateJWT(c.GetHeader("Authorization"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
