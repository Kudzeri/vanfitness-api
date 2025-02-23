package controllers

import (
	"context"
	"net/http"

	"github.com/Kudzeri/vanfitness-api/config"
	"github.com/Kudzeri/vanfitness-api/models"
	"github.com/gin-gonic/gin"
)

func MakeProfile(c *gin.Context) {
	var profile models.Profile
	if err:=c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid input"})
		return
	}

	_, err := config.ProfileCollection.InsertOne(context.TODO(), profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create profile"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}
