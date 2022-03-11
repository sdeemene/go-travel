package repository

import (
	"context"

	"github.com/stdeemene/go-travel/database"
	"github.com/stdeemene/go-travel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var addressCollection = database.AddressCollection()

func SaveAddress(ctx context.Context, doc models.Address) (interface{}, error) {
	result, err := addressCollection.InsertOne(ctx, doc)
	return result.InsertedID, err
}

func FindAddressById(ctx context.Context, filter primitive.M) (models.Address, error) {
	var address models.Address
	err := addressCollection.FindOne(ctx, filter).Decode(&address)
	return address, err
}

func UpdateAddress(ctx context.Context, filter primitive.M, address models.Address) (models.Address, error) {
	err := addressCollection.FindOneAndUpdate(ctx, filter, bson.M{"$set": address}, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&address)
	return address, err
}

func DeleteAddress(ctx context.Context, filter primitive.M) (int64, error) {
	res, err := addressCollection.DeleteOne(ctx, filter)
	return res.DeletedCount, err
}

func FindAllAddresses(ctx context.Context) ([]models.Address, error) {
	var addresses []models.Address
	findOptions := options.Find()
	findOptions.SetLimit(100)

	cursor, err := addressCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return addresses, err
	}
	// cursor.All(ctx, addresses)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		address := models.Address{}
		err := cursor.Decode(&address)
		if err != nil {
			return addresses, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func SearchAllAddresses(ctx context.Context, filter interface{}) ([]models.Address, error) {
	addresses := []models.Address{}
	cursor, err := addressCollection.Find(ctx, filter)

	if err != nil {
		return addresses, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		address := models.Address{}
		err := cursor.Decode(&address)
		if err != nil {
			return addresses, err
		}
		addresses = append(addresses, address)
	}

	return addresses, nil

}
