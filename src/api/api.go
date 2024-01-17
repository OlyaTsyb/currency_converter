package api

import (
	"encoding/json"
	"example/web-service-gin/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/url"
	"time"
)

func GetRatesFromAPI(fromCurrency, date string) ([]models.Exchange, error) {
	apiURL := fmt.Sprintf("https://www.floatrates.com/daily/%s.json", url.QueryEscape(fromCurrency))
	response, err := http.Get(apiURL)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status: %s", response.Status)
	}

	var data map[string]models.Rate

	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	timestamp := primitive.NewDateTimeFromTime(time.Now())
	var exchanges []models.Exchange

	exchange := models.Exchange{
		Rates:     make(map[string]models.Rate),
		Currency:  fromCurrency,
		Timestamp: timestamp,
	}

	for key, value := range data {
		value.Date = date
		exchange.Rates[key] = value
	}

	exchanges = append(exchanges, exchange)

	return exchanges, nil
}
