package mongo

import (
	"context"
	"example/web-service-gin/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27020")
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

func InsertRate(rate models.Rate) {
	rate.Timestamp = primitive.NewDateTimeFromTime(time.Now())
	collection := client.Database("mongodb").Collection("rates")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, rate)
	if err != nil {
		log.Fatal(err)
	}

}

func GetRateByCurrencyCode(code string) (float64, error) {
	collection := client.Database("mongodb").Collection("rates")
	filter := bson.M{"currency.code": code}

	var result models.Rate
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return 0, err
	}

	fmt.Printf("Rate: %f\n", result.Currency.Rate)
	return result.Currency.Rate, nil
}
func ClearRatesTable() error {
	collection := client.Database("mongodb").Collection("rates")

	_, err := collection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Rates table cleared")
	return nil
}
