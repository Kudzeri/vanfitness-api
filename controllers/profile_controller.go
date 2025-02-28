package controllers

import (
	"context"
	"net/http"

	"github.com/Kudzeri/vanfitness-api/models"
	"github.com/Kudzeri/vanfitness-api/repositories"
	"github.com/gin-gonic/gin"
)

// Создание профиля
func MakeProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.TODO()
	user, err := repositories.GetUserByUsername(ctx, username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if _, err := repositories.GetProfileByUserID(ctx, user.ID); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Profile already exists"})
		return
	}

	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Установка дефолтных значений
	profile.Prefix = "Fresh Boy"
	profile.Level = "1"
	profile.UserID = user.ID

	if err := repositories.CreateProfile(ctx, profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully"})
}

// Получение профиля
func GetProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.TODO()
	user, err := repositories.GetUserByUsername(ctx, username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	profile, err := repositories.GetProfileByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// Обновление профиля
func UpdateProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx := context.TODO()
	user, err := repositories.GetUserByUsername(ctx, username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Получаем текущий профиль
	currentProfile, err := repositories.GetProfileByUserID(ctx, user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Декодируем входящие данные
	var inputProfile models.Profile
	if err := c.ShouldBindJSON(&inputProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Обновляем только переданные параметры, оставляя остальные без изменений
	updatedProfile := models.Profile{
		UserID: user.ID,
		Height: ifEmpty(inputProfile.Height, currentProfile.Height),
		Weight: ifEmpty(inputProfile.Weight, currentProfile.Weight),
		Age:    ifEmpty(inputProfile.Age, currentProfile.Age),
		Sex:    ifEmpty(inputProfile.Sex, currentProfile.Sex),
	}

	if err := repositories.UpdateProfile(ctx, user.ID, updatedProfile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// Функция для проверки пустых строк
func ifEmpty(newValue, oldValue string) string {
	if newValue == "" {
		return oldValue
	}
	return newValue
}