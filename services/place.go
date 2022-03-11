package services

import (
	"context"
	"fmt"

	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type PlaceInterface interface {
// 	CreatePlaceDoc(models.Place) (models.Place, error)

// 	GetPlaceDocById(string) (models.Place, error)

// 	UpdatePlaceDoc(string, interface{}) (models.PlaceUpdated, error)

// 	DeletePlaceDoc(string) (models.PlaceDeleted, error)

// 	FindAllPlaces() ([]models.Place, error)

// 	SearchAllPlaces(interface{}) ([]models.Place, error)
// }

// var placeRepository repository.PlaceRepository

func CreatePlaceDoc(ctx context.Context, payload *models.PlaceReq) (*models.Place, error) {
	place := &models.Place{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Email:       payload.Email,
		Phone:       payload.Phone,
	}

	place.Initialize()

	address, err := GetAddressDocById(ctx, payload.AddressID)
	if err != nil {
		return nil, err
	}
	place.Address = &address

	result, err := repository.SavePlace(ctx, place)
	if err != nil {
		return place, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("Record Created: ", place)
	return GetPlaceDocById(ctx, id)
}

func GetPlaceDocById(ctx context.Context, id string) (*models.Place, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	place, err := repository.FindPlaceById(ctx, filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("Record Found: ", place)
	return &place, nil
}

func UpdatePlaceDoc(ctx context.Context, id string, place models.Place) (models.PlaceUpdated, error) {
	result := models.PlaceUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	oldplace, err := GetPlaceDocById(ctx, id)
	if err != nil {
		fmt.Println(oldplace)
		return result, err
	}

	res, err := repository.UpdatePlace(ctx, filter, place)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil
}

func DeletePlaceDoc(ctx context.Context, id string) (models.PlaceDeleted, error) {
	result := models.PlaceDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}
	place, err := GetPlaceDocById(ctx, id)
	if err != nil {
		fmt.Println(place)
		return result, err
	}

	res, err := repository.DeletePlace(ctx, filter)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	fmt.Println("Deleted a single document: ", result.DeletedCount)
	return result, nil
}

func FindAllPlaces(ctx context.Context) ([]models.Place, error) {
	places, err := repository.FindPlaces(ctx)
	if err != nil {
		return places, err
	}
	return places, err
}

func SearchAllPlaces(ctx context.Context, filter interface{}) ([]models.Place, error) {
	if filter == nil {
		filter = bson.M{}
	}

	places, err := repository.SearchPlaces(ctx, filter)

	if err != nil {
		return places, err
	}

	return places, nil

}
