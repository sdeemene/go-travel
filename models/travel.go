package models

import (
	"github.com/stdeemene/go-travel/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Travel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Duration    string             `bson:"duration,omitempty" json:"duration,omitempty"`
	TotalAmount float64            `bson:"totalAmount,omitempty" json:"totalAmount,omitempty"`
	User        *User              `bson:"user,omitempty" json:"user,omitempty"`
	Place       *Place             `bson:"place,omitempty" json:"place,omitempty"`
	TravelAT    string             `bson:"travelAt,omitempty" json:"travelAt,omitempty"`
	ReturnAT    string             `bson:"returnAt,omitempty" json:"returnAt,omitempty"`
	RateValue   int64              `bson:"rateValue,omitempty" json:"rateValue,omitempty"`
	CreatedAT   string             `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT   string             `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type TravelUpdated struct {
	ModifiedCount int64   `json:"modifiedCount"`
	Result        *Travel `json:"travelData"`
}

type TravelDeleted struct {
	DeletedCount int64 `json:"deletedCount"`
}

func (t *Travel) Initialize() error {
	t.CreatedAT = utility.CurrentDateTimeInString()
	t.UpdatedAT = utility.CurrentDateTimeInString()
	t.RateValue = 1
	return nil
}

type TravelReq struct {
	UserID   string `json:"userId"`
	PlaceID  string `json:"placeId"`
	TravelAT string `json:"travelDate"`
	ReturnAT string `json:"returnDate"`
}

type TravelRateReq struct {
	RateValue int `json:"rateValue"`
}
