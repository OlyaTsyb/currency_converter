package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Currency struct {
	Code        string  `json:"code"`
	AlphaCode   string  `json:"alphaCode"`
	NumericCode string  `json:"numericCode"`
	Name        string  `json:"name"`
	Rate        float64 `json:"rate"`
	Date        string  `json:"date"`
	InverseRate float64 `json:"inverseRate"`
}

type Rate struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Currency  Currency           `json:"currency"`
	Timestamp primitive.DateTime `json:"timestamp"`
}
