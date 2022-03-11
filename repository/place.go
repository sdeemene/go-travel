package repository

import (
	"context"

	"github.com/stdeemene/go-travel/database"
	"github.com/stdeemene/go-travel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type PlaceRepository interface {
// 	Save(models.Place) (interface{}, error)

// 	FindById(primitive.M) (models.Place, error)

// 	Update(primitive.M, interface{}) (int64, error)

// 	Delete(primitive.M) (int64, error)

// 	FindAll() ([]models.Place, error)

// 	Search(interface{}) ([]models.Place, error)
// }

var placeCollection = database.PlaceCollection()

func SavePlace(ctx context.Context, place *models.Place) (interface{}, error) {
	result, err := placeCollection.InsertOne(ctx, place)
	return result.InsertedID, err
}

func FindPlaceById(ctx context.Context, filter primitive.M) (models.Place, error) {
	var place models.Place
	err := placeCollection.FindOne(ctx, filter).Decode(&place)
	return place, err
}

func UpdatePlace(ctx context.Context, filter primitive.M, place models.Place) (models.Place, error) {
	err := placeCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": place}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&place)
	return place, err
}

func DeletePlace(ctx context.Context, filter primitive.M) (int64, error) {
	res, err := placeCollection.DeleteOne(ctx, filter)
	return res.DeletedCount, err
}

func FindPlaces(ctx context.Context) ([]models.Place, error) {
	var places []models.Place
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cursor, err := placeCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return places, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var place models.Place
		err := cursor.Decode(&place)
		if err != nil {
			return places, err
		}
		places = append(places, place)
	}

	return places, nil
}

func SearchPlaces(ctx context.Context, filter interface{}) ([]models.Place, error) {
	places := []models.Place{}
	cursor, err := placeCollection.Find(ctx, filter)

	if err != nil {
		return places, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		place := models.Place{}
		err := cursor.Decode(&place)
		if err != nil {
			return places, err
		}
		places = append(places, place)
	}

	return places, nil

}
