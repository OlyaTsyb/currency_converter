package handlers

import (
	"example/web-service-gin/mongo"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func HandleConvertRequest(c *gin.Context) {
	mongo.ConnectDB()

	fromCurrency := c.Query("from")
	toCurrency := c.Query("to")
	date := c.Query("date")
	amountStr := c.Query("amount")

	if fromCurrency == "---" && toCurrency == "---" {
		c.String(http.StatusBadRequest, "Please fill in the 'From' and 'To' fields")
		return
	}

	if fromCurrency == "---" {
		c.String(http.StatusBadRequest, "Please fill in the 'From' field")
		return
	}

	if toCurrency == "---" {
		c.String(http.StatusBadRequest, "Please fill in the 'To' field")
		return
	}

	if amountStr == "" {
		c.String(http.StatusBadRequest, "Please fill in the 'Amount' field")
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		c.String(http.StatusBadRequest, "Invalid 'amount' value")
		return
	}

	if amount == 0 {
		c.String(http.StatusBadRequest, "Please fill in a non-zero 'Amount' value")
		return
	}

	fmt.Printf("From: %s, To: %s, Date: %s, Amount: %f\n", fromCurrency, toCurrency, date, amount)

	targetRate, err := mongo.GetRateByCurrencyCode(fromCurrency, toCurrency, date)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "Failed to fetch rate from the database")
		return
	}

	result := targetRate * amount
	fmt.Printf("Result: %f\n", result)

	htmlResponse := fmt.Sprintf("%s: %f", toCurrency, result)

	c.String(http.StatusOK, htmlResponse)
}

func WelcomeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "welcomePage.html", nil)
}

func IndexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}
