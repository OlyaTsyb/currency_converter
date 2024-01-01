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

func GetRatesFromAPI(fromCurrency string) (map[string]models.Rate, error) {
	apiUrls := fmt.Sprintf("https://www.floatrates.com/daily/%s.json", url.QueryEscape(fromCurrency))
	response, err := http.Get(apiUrls)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK status: %s", response.Status)
	}

	var data map[string]models.Currency
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	rates := make(map[string]models.Rate)
	for key, value := range data {
		var rate models.Rate
		rate.Currency = value
		rate.Timestamp = primitive.NewDateTimeFromTime(time.Now())
		rates[key] = rate
	}

	return rates, nil
}
