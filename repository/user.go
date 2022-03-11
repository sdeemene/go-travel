package repository

import (
	"context"

	"github.com/stdeemene/go-travel/database"
	"github.com/stdeemene/go-travel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection = database.UserCollection()

func SaveUser(ctx context.Context, user *models.User) (interface{}, error) {
	result, err := userCollection.InsertOne(ctx, user)
	return result.InsertedID, err
}

func FindUser(ctx context.Context, filter primitive.M) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func UpdateUser(ctx context.Context, filter primitive.M, user models.User) (models.User, error) {
	err := userCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": user}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user)
	return user, err
}

func DeleteUser(ctx context.Context, filter primitive.M) (int64, error) {
	res, err := userCollection.DeleteOne(ctx, filter)
	return res.DeletedCount, err
}

func FindUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cursor, err := userCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return users, err
	}
	cursor.All(ctx, users)

	return users, nil
}

func SearchUsers(ctx context.Context, filter interface{}) ([]models.User, error) {
	users := []models.User{}
	cursor, err := userCollection.Find(ctx, filter)

	if err != nil {
		return users, err
	}
	cursor.All(ctx, users)
	return users, nil

}

func FindOne(ctx context.Context, filter primitive.M) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	return &user, err
}
