package services

import (
	"context"
	"fmt"

	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAddressDoc(ctx context.Context, doc models.Address) (models.Address, error) {
	var address models.Address
	result, err := repository.SaveAddress(ctx, doc)
	if err != nil {
		return address, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("Record Created: ", address)
	return GetAddressDocById(ctx, id)
}

func GetAddressDocById(ctx context.Context, id string) (models.Address, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	address, err := repository.FindAddressById(ctx, filter)
	if err != nil {
		return address, err
	}
	fmt.Println("Record Found: ", address)
	return address, nil
}

func UpdateAddressDoc(ctx context.Context, id string, address models.Address) (models.AddressUpdated, error) {
	result := models.AddressUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	existing, err := GetAddressDocById(ctx, id)
	if err != nil {
		fmt.Println(existing)
		return result, err
	}

	res, err := repository.UpdateAddress(ctx, filter, address)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil
}

func DeleteAddressDoc(ctx context.Context, id string) (models.AddressDeleted, error) {
	result := models.AddressDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	address, err := GetAddressDocById(ctx, id)
	if err != nil {
		fmt.Println(address)
		return result, err
	}

	res, err := repository.DeleteAddress(ctx, filter)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func FindAddresses(ctx context.Context) ([]models.Address, error) {
	addresses, err := repository.FindAllAddresses(ctx)
	if err != nil {
		return addresses, err
	}
	return addresses, err
}

func SearchAddresses(ctx context.Context, filter interface{}) ([]models.Address, error) {
	if filter == nil {
		filter = bson.M{}
	}

	addresses, err := repository.SearchAllAddresses(ctx, filter)

	if err != nil {
		return addresses, err
	}

	return addresses, nil

}
