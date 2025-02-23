package repositories

import (
	"context"
	"errors"

	"github.com/Kudzeri/vanfitness-api/config"
	"github.com/Kudzeri/vanfitness-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := config.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func UpdateUser(ctx context.Context, userID string, updatedUser models.User) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": updatedUser}

	result, err := config.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("could not update user")
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
