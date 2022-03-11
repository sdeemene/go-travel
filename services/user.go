package services

import (
	"context"
	"fmt"

	"github.com/stdeemene/go-travel/models"
	"github.com/stdeemene/go-travel/repository"
	"github.com/stdeemene/go-travel/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserDoc(ctx context.Context, payload *models.SignupReq) (*models.User, error) {
	requestUser := &models.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Gender:    payload.Gender,
		Email:     payload.Email,
		Password:  payload.Password,
		Phone:     payload.Phone,
	}

	requestUser.Initialize()
	result, err := repository.SaveUser(ctx, requestUser)
	if err != nil {
		return nil, err
	}
	id := result.(primitive.ObjectID).Hex()
	fmt.Println("Record Created: ", result)
	return GetUserDocById(ctx, id)
}

func GetUserDocById(ctx context.Context, id string) (*models.User, error) {
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	user, err := repository.FindUser(ctx, filter)
	if err != nil {
		return nil, err
	}

	fmt.Println("Record Found: ", user)
	return &user, nil
}

func UpdateUserDoc(ctx context.Context, id string, user models.User) (models.UserUpdated, error) {
	result := models.UserUpdated{
		ModifiedCount: 0,
	}
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, err
	}
	filter := bson.M{"_id": _id}

	res, err := repository.UpdateUser(ctx, filter, user)
	if err != nil {
		return result, err
	}

	result.ModifiedCount = 1
	result.Result = res
	fmt.Println("Record updated: ", result)
	return result, nil
}

func DeleteUserDoc(ctx context.Context, id string) (models.UserDeleted, error) {
	result := models.UserDeleted{
		DeletedCount: 0,
	}

	_id, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	if err != nil {
		return result, err
	}

	user, err := GetUserDocById(ctx, id)
	if err != nil {
		fmt.Println(user)
		return result, err
	}

	res, err := repository.DeleteUser(ctx, filter)
	if err != nil {
		return result, err
	}
	result.DeletedCount = res
	return result, nil
}

func FindAllUsers(ctx context.Context) ([]models.User, error) {
	users, err := repository.FindUsers(ctx)
	if err != nil {
		return nil, err
	}

	if IsEmpty(users) {
		return users, utility.NewError("no users found.")
	}
	return users, nil
}

func SearchAllUsers(ctx context.Context, filter interface{}) ([]models.User, error) {
	if filter == nil {
		filter = bson.M{}
	}

	users, err := repository.SearchUsers(ctx, filter)

	if err != nil {
		return nil, err
	}
	if IsEmpty(users) {
		return users, utility.NewError("no users found.")
	}
	return users, nil

}

func IsEmpty(users []models.User) bool {
	return len(users) == 0
}

func Login(ctx context.Context, credentials *models.LoginReq) (*models.User, error) {
	filter := bson.M{"email": credentials.Email}
	user, err := repository.FindOne(ctx, filter)

	if err != nil {
		return nil, utility.NewError("invalid email")
	}
	err = user.ComparePassword(credentials.Password)
	if err != nil {
		return nil, utility.NewError("invalid password")
	}
	return user, nil
}
