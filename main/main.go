package main

import (
	"example/web-service-gin/mongo"
	"example/web-service-gin/src/api"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

var listRates = []string{
	"mga", "srd", "syp", "mop", "bam", "nzd", "try", "ngn", "rsd", "nio", "sbd", "lak", "gtq", "nok", "qar",
	"czk", "byn", "ars", "stn", "bif", "aoa", "mvr", "ves", "bdt", "ron", "mdl", "crc", "bzd", "gnf", "hnl", "kes",
	"aed", "idr", "mxn", "amd", "pyg", "gyd", "rwf", "mzn", "ugx", "eur", "gbp", "jpy", "aud", "chf", "cad", "inr",
	"npr", "xaf", "kgs", "afn", "nad", "sdg", "top", "vuv", "brl", "lkr", "tnd", "vnd", "tmt", "svc", "xcd", "pgk",
	"bwp", "dkk", "pkr", "bgn", "rub", "gel", "mkd", "awg", "zmw", "khr", "sar", "pln", "kzt", "cop", "bbd", "djf",
	"all", "scr", "bhd", "egp", "krw", "dzd", "pab", "fjd", "cdf", "lsl", "tzs", "hkd", "mad", "zar", "iqd", "bob",
	"lrd", "ssp", "mru", "mnt", "kwd", "thb", "twd", "uzs", "etb", "ttd", "ghs", "cup", "omr", "ils", "pen",
	"tjs", "gmd", "cve", "mwk", "yer", "sek", "sgd", "huf", "uah", "usd",
}

func Task() {
	//timestamp := primitive.NewDateTimeFromTime(time.Now())
	date := time.Now().Format(time.RFC1123)
	for _, fromCurrency := range listRates {
		rates, err := api.GetRatesFromAPI(fromCurrency, date)
		fmt.Println("Rates", rates)
		if err != nil {
			log.Fatal(err)
			return
		}

		for _, exchange := range rates {
			exchange.Timestamp = primitive.NewDateTimeFromTime(time.Now())
			mongo.InsertRate(exchange)
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//gocron.Every(15).Minute().Do(main.Task)
	gocron.Every(1).Day().At("23:00").Do(Task)
	<-gocron.Start()

}
