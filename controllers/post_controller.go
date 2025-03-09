package controllers

import (
	"context"
	"net/http"

	"github.com/Kudzeri/vanfitness-api/models"
	"github.com/Kudzeri/vanfitness-api/repositories"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	ctx := context.TODO()

	//Получение всех постов
	posts, err := repositories.GetPosts(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	ctx := context.Background()

	//Проверка на наличие ID
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	//Получение поста по ID
	post, err := repositories.GetPostByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func CreatePost(c *gin.Context) {
	ctx := context.Background()

	//Получение ID пользователя из токена
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	//Получение данных из запроса
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.UserID = userID.(string)

	err := repositories.CreatePost(ctx, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func UpdatePost(c *gin.Context) {
	ctx := context.Background()

	// Получаем user_id из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Проверка на наличие ID поста
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	// Получаем пост из базы по ID
	post, err := repositories.GetPostByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Проверяем, является ли авторизованный пользователь владельцем поста
	if post.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to edit this post"})
		return
	}

	// Получаем данные из запроса
	var updatedPost models.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Обновляем только поля Title и Body (чтобы UserID не менялся)
	post.Title = updatedPost.Title
	post.Body = updatedPost.Body

	// Сохраняем изменения
	err = repositories.UpdatePost(ctx, id, post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	ctx := context.Background()

	// Получаем user_id из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Проверка на наличие ID поста
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	// Получаем пост из базы по ID
	post, err := repositories.GetPostByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Проверяем, является ли авторизованный пользователь владельцем поста
	if post.UserID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this post"})
		return
	}

	// Удаляем пост
	err = repositories.DeletePost(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
