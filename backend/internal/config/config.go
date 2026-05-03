package config

import (
	"os"
	"strconv"
)

type Config struct {
	QWeatherAPIKey  string
	QWeatherAPIHost string
	Port            int
}

func Load() *Config {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}
	return &Config{
		QWeatherAPIKey:  os.Getenv("QWEATHER_API_KEY"),
		QWeatherAPIHost: os.Getenv("QWEATHER_API_HOST"),
		Port:            port,
	}
}
