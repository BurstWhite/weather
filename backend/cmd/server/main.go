package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"weather-app/backend/internal/config"
	"weather-app/backend/internal/database"
	"weather-app/backend/internal/handler"
	"weather-app/backend/internal/qweather"
	"weather-app/backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(filepath.Join("..", ".env"))
	_ = godotenv.Load()

	cfg := config.Load()

	if cfg.QWeatherAPIKey == "" {
		log.Fatal("QWEATHER_API_KEY environment variable is required")
	}

	dbPath := filepath.Join(".", "weather.db")
	if err := database.Init(dbPath); err != nil {
		log.Fatalf("database init: %v", err)
	}
	defer database.Close()

	qClient := qweather.NewClient(cfg.QWeatherAPIKey, cfg.QWeatherAPIHost)
	weatherSvc := service.NewWeatherService(qClient)

	cityH := handler.NewCityHandler()
	weatherH := handler.NewWeatherHandler(weatherSvc)
	searchH := handler.NewSearchHandler(qClient)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		api.GET("/cities", cityH.ListCities)
		api.POST("/cities", cityH.AddCity)
		api.DELETE("/cities/:id", cityH.DeleteCity)
		api.GET("/search", searchH.SearchCity)
		api.GET("/weather/:id", weatherH.GetWeather)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server: %v", err)
		os.Exit(1)
	}
}
