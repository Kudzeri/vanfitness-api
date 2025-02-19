package controllers

import (
	"context"
	"net/http"

	"github.com/Kudzeri/vanfitness-api/config"
	"github.com/Kudzeri/vanfitness-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(c *gin.Context) {
	username, _ := c.Get("username")
	var user models.User
	err := config.UserCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
