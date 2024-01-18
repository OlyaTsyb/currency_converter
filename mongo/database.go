package mongo

import (
	"context"
	"errors"
	"example/web-service-gin/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

func init() {
	ConnectDB()
}

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

func InsertRate(exchange models.Exchange) {
	exchange.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	collection := client.Database("mongodb").Collection("rates")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, exchange)
	if err != nil {
		log.Fatal(err)
	}

}

func GetRateByCurrencyCode(code, toCurrency, date string) (float64, error) {
	collection := client.Database("mongodb").Collection("rates")

	dateObj, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, err
	}

	startOfDay := time.Date(dateObj.Year(), dateObj.Month(), dateObj.Day(), 0, 0, 0, 0, dateObj.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond)

	filter := bson.M{
		"currency": code,
		"timestamp": bson.M{
			"$gte": primitive.NewDateTimeFromTime(startOfDay),
			"$lt":  primitive.NewDateTimeFromTime(endOfDay),
		},
	}

	var result models.Exchange
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return 0, err
	}

	rates := result.Rates

	if rate, ok := rates[toCurrency]; ok {
		return rate.Rate, nil
	}

	return 0, errors.New("no rate found for the specified toCurrency")
}

func GetRateByCurrencyCodeAndDate(code, date string) (map[string]models.Rate, error) {
	collection := client.Database("mongodb").Collection("rates")
	dateObj, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}

	startOfDay := time.Date(dateObj.Year(), dateObj.Month(), dateObj.Day(), 0, 0, 0, 0, dateObj.Location())
	endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond)

	filter := bson.M{
		"currency": code,
		"timestamp": bson.M{
			"$gte": primitive.NewDateTimeFromTime(startOfDay),
			"$lt":  primitive.NewDateTimeFromTime(endOfDay),
		},
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var result models.Exchange
	rates := make(map[string]models.Rate)

	for cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		rates = result.Rates
	}

	return rates, nil
}
