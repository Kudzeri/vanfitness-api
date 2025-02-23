package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Kudzeri/vanfitness-api/utils"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            c.Abort()
            return
        }

        if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
            authHeader = authHeader[7:]
        }

        username, err := utils.ValidateJWT(authHeader)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        c.Set("username", username)
        c.Next()
    }
}

