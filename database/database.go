package database

import (
	"context"
	"log"

	"github.com/stdeemene/go-travel/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context, config config.Configuration) *mongo.Database {
	var clientOption *options.ClientOptions
	if config.Environment.Profile == "dev" {
		clientOption = options.Client().ApplyURI(config.Database.LocalUri)
	} else {
		clientOption = options.Client().ApplyURI(config.Database.RemoteUri)
	}

	client, err := mongo.Connect(context.Background(), clientOption)

	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}

	return client.Database(config.Database.DatabaseName)

}

func PlaceCollection() *mongo.Collection {
	config := config.GetConfiguration()
	database := ConnectDB(context.TODO(), config)
	placeCollection := database.Collection("place")
	return placeCollection
}

func AddressCollection() *mongo.Collection {
	config := config.GetConfiguration()
	database := ConnectDB(context.TODO(), config)
	addressCollection := database.Collection("address")
	return addressCollection
}

func UserCollection() *mongo.Collection {
	config := config.GetConfiguration()
	database := ConnectDB(context.TODO(), config)
	userCollection := database.Collection("user")
	return userCollection
}

func TravelCollection() *mongo.Collection {
	config := config.GetConfiguration()
	database := ConnectDB(context.TODO(), config)
	travelCollection := database.Collection("travel")
	return travelCollection
}
