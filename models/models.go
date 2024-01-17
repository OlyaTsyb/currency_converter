package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rate struct {
	Code        string  `json:"code"`
	AlphaCode   string  `json:"alphaCode"`
	NumericCode string  `json:"numericCode"`
	Name        string  `json:"name"`
	Rate        float64 `json:"rate"`
	Date        string  `json:"date"`
	InverseRate float64 `json:"inverseRate"`
}

type Exchange struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rates     map[string]Rate    `json:"rates"`
	Currency  string             `json:"currency"`
	Timestamp primitive.DateTime `json:"timestamp"`
}
