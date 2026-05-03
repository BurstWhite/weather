package model

import "time"

type City struct {
	ID         int64   `json:"id"`
	LocationID string  `json:"locationId"`
	Name       string  `json:"name"`
	Adm1       string  `json:"adm1"`
	Adm2       string  `json:"adm2"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
}

type WeatherDaily struct {
	ID           int64     `json:"id"`
	CityID       int64     `json:"cityId"`
	FxDate       string    `json:"fxDate"`
	TempMax      int       `json:"tempMax"`
	TempMin      int       `json:"tempMin"`
	TextDay      string    `json:"textDay"`
	TextNight    string    `json:"textNight"`
	IconDay      string    `json:"iconDay"`
	IconNight    string    `json:"iconNight"`
	Humidity     int       `json:"humidity"`
	WindDirDay   string    `json:"windDirDay"`
	WindScaleDay string    `json:"windScaleDay"`
	Precip       float64   `json:"precip"`
	UVIndex      int       `json:"uvIndex"`
	Pressure     float64   `json:"pressure"`
	Sunrise      string    `json:"sunrise"`
	Sunset       string    `json:"sunset"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type AirQuality struct {
	ID               int64     `json:"id"`
	CityID           int64     `json:"cityId"`
	Aqi              int       `json:"aqi"`
	Category         string    `json:"category"`
	Level            string    `json:"level"`
	PrimaryPollutant string    `json:"primaryPollutant"`
	Pm2p5            float64   `json:"pm2p5"`
	Pm10             float64   `json:"pm10"`
	No2              float64   `json:"no2"`
	O3               float64   `json:"o3"`
	Co               float64   `json:"co"`
	So2              float64   `json:"so2"`
	HealthEffect     string    `json:"healthEffect"`
	HealthAdvice     string    `json:"healthAdvice"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type WeatherResponse struct {
	City       City            `json:"city"`
	Weather    []WeatherDaily  `json:"weather"`
	AirQuality *AirQuality     `json:"airQuality,omitempty"`
}
