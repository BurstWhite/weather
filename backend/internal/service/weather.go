package service

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"weather-app/backend/internal/database"
	"weather-app/backend/internal/model"
	"weather-app/backend/internal/qweather"
)

const cacheTTL = 1 * time.Hour

type WeatherService struct {
	q *qweather.Client
}

func NewWeatherService(q *qweather.Client) *WeatherService {
	return &WeatherService{q: q}
}

func (s *WeatherService) GetWeather(cityID int64, lat, lon float64, locationID string) (*model.WeatherResponse, error) {
	city := model.City{}
	row := database.DB.QueryRow("SELECT id, location_id, name, adm1, adm2, lat, lon FROM cities WHERE id = ?", cityID)
	if err := row.Scan(&city.ID, &city.LocationID, &city.Name, &city.Adm1, &city.Adm2, &city.Lat, &city.Lon); err != nil {
		return nil, fmt.Errorf("city not found: %w", err)
	}

	weather, err := s.getCachedWeather(city.ID, city.LocationID)
	if err != nil {
		return nil, err
	}

	air, err := s.getCachedAirQuality(city.ID, city.Lat, city.Lon)
	if err != nil {
		return nil, err
	}

	return &model.WeatherResponse{
		City:       city,
		Weather:    weather,
		AirQuality: air,
	}, nil
}

func (s *WeatherService) getCachedWeather(cityID int64, locationID string) ([]model.WeatherDaily, error) {
	rows, err := database.DB.Query(
		"SELECT fx_date, temp_max, temp_min, text_day, text_night, icon_day, icon_night, humidity, wind_dir_day, wind_scale_day, precip, uv_index, pressure, sunrise, sunset, updated_at FROM weather_cache WHERE city_id = ? ORDER BY fx_date",
		cityID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cached []model.WeatherDaily
	for rows.Next() {
		var w model.WeatherDaily
		w.CityID = cityID
		if err := rows.Scan(&w.FxDate, &w.TempMax, &w.TempMin, &w.TextDay, &w.TextNight, &w.IconDay, &w.IconNight, &w.Humidity, &w.WindDirDay, &w.WindScaleDay, &w.Precip, &w.UVIndex, &w.Pressure, &w.Sunrise, &w.Sunset, &w.UpdatedAt); err != nil {
			return nil, err
		}
		cached = append(cached, w)
	}

	if len(cached) == 7 {
		oldest := cached[0].UpdatedAt
		if time.Since(oldest) < cacheTTL {
			return cached, nil
		}
	}

	days, err := s.q.GetWeather7Days(locationID)
	if err != nil {
		if len(cached) > 0 {
			return cached, nil
		}
		return nil, err
	}

	tx, err := database.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	tx.Exec("DELETE FROM weather_cache WHERE city_id = ?", cityID)

	stmt, err := tx.Prepare(`INSERT INTO weather_cache(city_id, fx_date, temp_max, temp_min, text_day, text_night, icon_day, icon_night, humidity, wind_dir_day, wind_scale_day, precip, uv_index, pressure, sunrise, sunset, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, d := range days {
		tMax, _ := strconv.Atoi(d.TempMax)
		tMin, _ := strconv.Atoi(d.TempMin)
		hum, _ := strconv.Atoi(d.Humidity)
		uv, _ := strconv.Atoi(d.UVIndex)
		precip, _ := strconv.ParseFloat(d.Precip, 64)
		pressure, _ := strconv.ParseFloat(d.Pressure, 64)

		_, err := stmt.Exec(cityID, d.FxDate, tMax, tMin, d.TextDay, d.TextNight, d.IconDay, d.IconNight, hum, d.WindDirDay, d.WindScaleDay, precip, uv, pressure, d.Sunrise, d.Sunset, time.Now())
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return s.getCachedWeather(cityID, locationID)
}

func (s *WeatherService) getCachedAirQuality(cityID int64, lat, lon float64) (*model.AirQuality, error) {
	row := database.DB.QueryRow("SELECT aqi, category, level, primary_pollutant, pm2p5, pm10, no2, o3, co, so2, health_effect, health_advice, updated_at FROM air_cache WHERE city_id = ? ORDER BY updated_at DESC LIMIT 1", cityID)
	air := &model.AirQuality{CityID: cityID}
	err := row.Scan(&air.Aqi, &air.Category, &air.Level, &air.PrimaryPollutant, &air.Pm2p5, &air.Pm10, &air.No2, &air.O3, &air.Co, &air.So2, &air.HealthEffect, &air.HealthAdvice, &air.UpdatedAt)
	if err == nil && time.Since(air.UpdatedAt) < cacheTTL {
		return air, nil
	}

	result, err := s.q.GetAirQuality(lat, lon)
	if err != nil {
		if air.Aqi > 0 {
			return air, nil
		}
		return nil, err
	}

	air = &model.AirQuality{CityID: cityID}
	for _, idx := range result.Indexes {
		if idx.Code == "qaqi" || idx.Code == "cn-epa" {
			air.Aqi = int(math.Round(idx.Aqi))
			air.Category = idx.Category
			air.Level = idx.Level
			if idx.PrimaryPollutant.Code != "" {
				air.PrimaryPollutant = idx.PrimaryPollutant.Name
			}
			air.HealthEffect = idx.Health.Effect
			air.HealthAdvice = idx.Health.Advice.GeneralPopulation
		}
	}
	for _, p := range result.Pollutants {
		switch p.Code {
		case "pm2p5":
			air.Pm2p5 = p.Concentration.Value
		case "pm10":
			air.Pm10 = p.Concentration.Value
		case "no2":
			air.No2 = p.Concentration.Value
		case "o3":
			air.O3 = p.Concentration.Value
		case "co":
			air.Co = p.Concentration.Value
		case "so2":
			air.So2 = p.Concentration.Value
		}
	}

	database.DB.Exec("DELETE FROM air_cache WHERE city_id = ?", cityID)
	database.DB.Exec(`INSERT INTO air_cache(city_id, aqi, category, level, primary_pollutant, pm2p5, pm10, no2, o3, co, so2, health_effect, health_advice, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		cityID, air.Aqi, air.Category, air.Level, air.PrimaryPollutant, air.Pm2p5, air.Pm10, air.No2, air.O3, air.Co, air.So2, air.HealthEffect, air.HealthAdvice, time.Now())

	return air, nil
}
