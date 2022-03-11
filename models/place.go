package models

import (
	"github.com/stdeemene/go-travel/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Place struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64            `bson:"price,omitempty" json:"price,omitempty"`
	IsAvailable bool               `bson:"isAvailable,omitempty" json:"isAvailable,omitempty"`
	Phone       string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Address     *Address           `bson:"address,omitempty" json:"address,omitempty"`
	CreatedAT   string             `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT   string             `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type PlaceUpdated struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        Place
}

type PlaceDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}

func (p *Place) Initialize() error {
	p.CreatedAT = utility.CurrentDateTimeInString()
	p.UpdatedAT = utility.CurrentDateTimeInString()
	p.IsAvailable = true
	return nil
}

type PlaceReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	AddressID   string  `json:"addressId"`
}
