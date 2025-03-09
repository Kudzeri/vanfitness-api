package repositories

import (
	"context"

	"github.com/Kudzeri/vanfitness-api/models"
	"github.com/Kudzeri/vanfitness-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
)

func GetPosts(ctx context.Context) (*[]models.Post, error) {
	posts := []models.Post{}

	// Получаем курсор и ошибку
	cursor, err := config.PostCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // Закрываем курсор после использования

	// Декодируем найденные документы в массив `posts`
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post

	// Ищем пост по ID
	err := config.PostCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func CreatePost(ctx context.Context, post *models.Post) error {
	// Вставляем пост в коллекцию
	_, err := config.PostCollection.InsertOne(ctx, post)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(ctx context.Context, id string, post *models.Post) error {
	// Преобразуем id из строки в ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid post ID format: %v", err)
	}

	// Обновляем только определённые поля, а не всю структуру
	update := bson.M{"$set": bson.M{
		"title": post.Title,
		"body":  post.Body,
	}}

	// Обновляем пост по ID
	res, err := config.PostCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	// Проверяем, был ли обновлён хотя бы один документ
	if res.MatchedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
}

func DeletePost(ctx context.Context, id string) error {
	// Преобразуем id из строки в ObjectID
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid post ID format: %v", err)
	}

	// Удаляем пост по ID
	res, err := config.PostCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	// Проверяем, был ли удалён хотя бы один документ
	if res.DeletedCount == 0 {
		return fmt.Errorf("post not found")
	}

	return nil
	}
