package handlers

import (
	"example/web-service-gin/mongo"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
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

func CurrencyHistoryHandler(c *gin.Context) {
	mongo.ConnectDB()

	code := c.Query("from")

	date := c.Query("date")
	if code == "" || code == "0" {
		code = "usd"
	}

	if date == "" {
		date = time.Now().Format("2006-01-02")
	} else {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			log.Println("Failed to parse date:", err)
			date = time.Now().Format("2006-01-02")
		} else {
			fmt.Println("Parsed date:", parsedDate)
		}
	}

	fmt.Printf("From: %s, Date: %s\n", code, date)

	targetRate, err := mongo.GetRateByCurrencyCodeAndDate(code, date)
	if err != nil {
		log.Fatal(err)
		c.String(http.StatusInternalServerError, "Failed to fetch rate from the database")
		return
	}

	data := gin.H{
		"rates":        targetRate,
		"fromCurrency": code,
		"selectedDate": date,
	}

	c.HTML(http.StatusOK, "chart.html", data)
}

func IndexHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", nil)
}
