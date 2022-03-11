package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hako/durafmt"
	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/repository"
	"github.com/stdeemene/go-travel/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTravelDoc(ctx context.Context, payload *models.TravelReq) (*models.Travel, error) {

	travelDate, err := utility.ConvertToDate(payload.TravelAT)
	if err != nil {
		return nil, err
	}

	returnDate, err := utility.ConvertToDate(payload.ReturnAT)
	if err != nil {
		return nil, err
	}

	travel := &models.Travel{
		TravelAT: payload.TravelAT,
		ReturnAT: payload.ReturnAT,
	}
	travel.Initialize()

	user, err := GetUserDocById(ctx, payload.UserID)
	if err != nil {
		return nil, err
	}
	travel.User = user

	interval := returnDate.Sub(travelDate)
	intervalInDays := interval.Hours() / 24
	duration, err := durafmt.ParseString(interval.String())
	if err != nil {
		return nil, err
	}
	travel.Duration = duration.String()

	place, err := GetPlaceDocById(ctx, payload.PlaceID)
	if err != nil {
		return nil, err
	}
	travel.Place = place

	calculatedAmount := intervalInDays * place.Price
	amount := strconv.FormatFloat(calculatedAmount, 'f', 2, 64)
	totalAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}
	travel.TotalAmount = totalAmount
	result, err := repository.SaveTravel(ctx, travel)
	if err != nil {
		return travel, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("Record Created: ", travel)
	return GetTravelDocById(ctx, id)
}

func GetTravelDocById(ctx context.Context, id string) (*models.Travel, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	travel, err := repository.FindTravelById(ctx, filter)
	if err != nil {
		return nil, err
	}
	fmt.Println("Record Found: ", travel)
	return &travel, nil
}

func UpdateTravelDoc(ctx context.Context, id string, travel models.Travel) (models.TravelUpdated, error) {
	result := models.TravelUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	existing, err := GetTravelDocById(ctx, id)
	if err != nil {
		fmt.Println(existing)
		return result, err
	}

	res, err := repository.UpdateTravel(ctx, filter, travel)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = &res
	fmt.Println("Record updated: ", result)
	return result, nil
}

func DeleteTravelDoc(ctx context.Context, id string) (models.TravelDeleted, error) {
	result := models.TravelDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	travel, err := GetTravelDocById(ctx, id)
	if err != nil {
		fmt.Println(travel)
		return result, err
	}

	res, err := repository.DeleteTravel(ctx, filter)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func FindTravels(ctx context.Context) ([]models.Travel, error) {
	travels, err := repository.FindAllTravels(ctx)
	if err != nil {
		return travels, err
	}
	return travels, err
}

func SearchTravels(ctx context.Context, filter interface{}) ([]models.Travel, error) {
	if filter == nil {
		filter = bson.M{}
	}

	travels, err := repository.SearchAllTravels(ctx, filter)

	if err != nil {
		return travels, err
	}

	return travels, nil

}

func GetUserTravels(ctx context.Context, id string) ([]models.Travel, error) {
	user, err := GetUserDocById(ctx, id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"user": user}
	travels, err := repository.FindAllUserTravels(ctx, filter)
	if err != nil {
		return nil, err
	}
	return travels, nil
}

func UpdateTravelRateValue(ctx context.Context, id string, rateValue *models.TravelRateReq) (models.TravelUpdated, error) {

	rating := &models.TravelRateReq{
		RateValue: rateValue.RateValue,
	}

	fmt.Println("rating.RateValue 1  ", rateValue.RateValue)
	fmt.Println("rating.RateValue", rating.RateValue)
	result := models.TravelUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	_, err = GetTravelDocById(ctx, id)
	if err != nil {
		return result, err
	}
	update := bson.M{"$set": bson.M{"rateValue": rating.RateValue}}
	value, err := repository.UpdateTravelRateValue(ctx, filter, update)

	if err != nil {
		return result, err
	}
	travel, err := GetTravelDocById(ctx, id)
	if err != nil {
		return result, err
	}
	result.ModifiedCount = value
	result.Result = travel
	return result, nil
}
