package repository

import (
	"context"

	"github.com/stdeemene/go-travel/database"
	"github.com/stdeemene/go-travel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var travelCollection = database.TravelCollection()

func SaveTravel(ctx context.Context, travel *models.Travel) (interface{}, error) {
	result, err := travelCollection.InsertOne(ctx, travel)
	return result.InsertedID, err
}

func FindTravelById(ctx context.Context, filter primitive.M) (models.Travel, error) {
	var travel models.Travel
	err := travelCollection.FindOne(ctx, filter).Decode(&travel)
	return travel, err
}

func UpdateTravel(ctx context.Context, filter primitive.M, travel models.Travel) (models.Travel, error) {
	err := travelCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": travel}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&travel)
	return travel, err
}

func DeleteTravel(ctx context.Context, filter primitive.M) (int64, error) {
	res, err := travelCollection.DeleteOne(ctx, filter)
	return res.DeletedCount, err
}

func FindAllTravels(ctx context.Context) ([]models.Travel, error) {
	var travels []models.Travel
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cursor, err := travelCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return travels, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var travel models.Travel
		err := cursor.Decode(&travel)
		if err != nil {
			return travels, err
		}
		travels = append(travels, travel)
	}

	return travels, nil
}

func SearchAllTravels(ctx context.Context, filter interface{}) ([]models.Travel, error) {
	travels := []models.Travel{}
	cursor, err := travelCollection.Find(ctx, filter)

	if err != nil {
		return travels, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		travel := models.Travel{}
		err := cursor.Decode(&travel)
		if err != nil {
			return travels, err
		}
		travels = append(travels, travel)
	}

	return travels, nil

}

func FindAllUserTravels(ctx context.Context, filter interface{}) ([]models.Travel, error) {
	var travels []models.Travel
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cursor, err := travelCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return travels, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var travel models.Travel
		err := cursor.Decode(&travel)
		if err != nil {
			return travels, err
		}
		travels = append(travels, travel)
	}

	return travels, nil
}

func UpdateTravelRateValue(ctx context.Context, filter primitive.M, update primitive.M) (int64, error) {
	// travel := models.Travel{}
	result, err := travelCollection.UpdateOne(ctx, filter, update)
	return result.ModifiedCount, err
	//
	// err := userCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": bson.M{"rateValue": update}}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&travel)
	// fmt.Println("err.Error()", err.Error())
	// fmt.Println("travel", travel)
	// return travel, err
}
