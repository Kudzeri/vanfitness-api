package repositories

import (
	"context"
	"errors"

	"github.com/Kudzeri/vanfitness-api/config"
	"github.com/Kudzeri/vanfitness-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetProfileByUserID(ctx context.Context, userID string) (*models.Profile, error) {
	var profile models.Profile
	err := config.ProfileCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&profile)
	if err != nil {
		return nil, errors.New("profile not found")
	}
	return &profile, nil
}

func CreateProfile(ctx context.Context, profile models.Profile) error {
	_, err := config.ProfileCollection.InsertOne(ctx, profile)
	if err != nil {
		return errors.New("could not create profile")
	}
	return nil
}

func UpdateProfile(ctx context.Context, userID string, updatedProfile models.Profile) error {
	filter := bson.M{"user_id": userID}
	update := bson.M{"$set": updatedProfile}

	result, err := config.ProfileCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("could not update profile")
	}
	if result.MatchedCount == 0 {
		return errors.New("profile not found")
	}
	return nil
}