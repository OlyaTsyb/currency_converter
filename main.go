package main

import (
	"example/web-service-gin/src/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	templatesPath := filepath.Join(getEnv("TEMPLATES_PATH", ""), "*")

	staticPath := filepath.Join(getEnv("STATIC_PATH", ""), "")

	r.Static("/static", staticPath)

	r.LoadHTMLGlob(templatesPath)

	r.GET("/", handlers.WelcomeHandler)

	r.GET("/index", handlers.IndexHandler)

	r.GET("/convert", handlers.HandleConvertRequest)

	r.GET("/history", handlers.CurrencyHistoryHandler)

	return r
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := setupRouter()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
