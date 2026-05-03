package qweather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL     = "https://api.qweather.com"
	geoBaseURL  = "https://geoapi.qweather.com"
	airBaseURL  = "https://airquality.api.qweather.com"
)

type Client struct {
	apiKey  string
	apiHost string
	http    *http.Client
}

func NewClient(apiKey, apiHost string) *Client {
	return &Client{
		apiKey:  apiKey,
		apiHost: apiHost,
		http:    &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *Client) baseURL() string {
	if c.apiHost != "" {
		return "https://" + c.apiHost
	}
	return baseURL
}

func (c *Client) geoBaseURL() string {
	if c.apiHost != "" {
		return "https://" + c.apiHost
	}
	return geoBaseURL
}

func (c *Client) airBaseURL() string {
	if c.apiHost != "" {
		return "https://" + c.apiHost
	}
	return airBaseURL
}

func (c *Client) doGet(base, path string, params map[string]string) ([]byte, error) {
	u, _ := url.Parse(base + path)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("X-QW-Api-Key", c.apiKey)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}
	return body, nil
}

type CityLookupResult struct {
	Code     string          `json:"code"`
	Location []CityLocation  `json:"location"`
}

type CityLocation struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm1      string `json:"adm1"`
	Adm2      string `json:"adm2"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UTCOffset string `json:"utcOffset"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}

func (c *Client) SearchCity(query string) ([]CityLocation, error) {
	geoPath := "/v2/city/lookup"
	if c.apiHost != "" {
		geoPath = "/geo/v2/city/lookup"
	}
	body, err := c.doGet(c.geoBaseURL(), geoPath, map[string]string{
		"location": query,
	})
	if err != nil {
		return nil, err
	}
	var result CityLookupResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.Code != "200" {
		return nil, fmt.Errorf("qweather error code: %s", result.Code)
	}
	return result.Location, nil
}

type WeatherDailyResult struct {
	Code       string          `json:"code"`
	UpdateTime string          `json:"updateTime"`
	Daily      []WeatherDay    `json:"daily"`
}

type WeatherDay struct {
	FxDate        string `json:"fxDate"`
	Sunrise       string `json:"sunrise"`
	Sunset        string `json:"sunset"`
	TempMax       string `json:"tempMax"`
	TempMin       string `json:"tempMin"`
	IconDay       string `json:"iconDay"`
	TextDay       string `json:"textDay"`
	IconNight     string `json:"iconNight"`
	TextNight     string `json:"textNight"`
	WindDirDay    string `json:"windDirDay"`
	WindScaleDay  string `json:"windScaleDay"`
	Humidity      string `json:"humidity"`
	Precip        string `json:"precip"`
	Pressure      string `json:"pressure"`
	Vis           string `json:"vis"`
	Cloud         string `json:"cloud"`
	UVIndex       string `json:"uvIndex"`
}

func (c *Client) GetWeather7Days(locationID string) ([]WeatherDay, error) {
	body, err := c.doGet(c.baseURL(), "/v7/weather/7d", map[string]string{
		"location": locationID,
	})
	if err != nil {
		return nil, err
	}
	var result WeatherDailyResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.Code != "200" {
		return nil, fmt.Errorf("qweather error code: %s", result.Code)
	}
	return result.Daily, nil
}

type AirQualityResult struct {
	Indexes    []AirIndex     `json:"indexes"`
	Pollutants []AirPollutant `json:"pollutants"`
}

type AirIndex struct {
	Code             string          `json:"code"`
	Name             string          `json:"name"`
	Aqi              float64         `json:"aqi"`
	AqiDisplay       string          `json:"aqiDisplay"`
	Level            string          `json:"level"`
	Category         string          `json:"category"`
	PrimaryPollutant AirPollutantRef `json:"primaryPollutant"`
	Health           HealthInfo      `json:"health"`
}

type HealthInfo struct {
	Effect  string       `json:"effect"`
	Advice AdviceDetail `json:"advice"`
}

type AdviceDetail struct {
	GeneralPopulation  string `json:"generalPopulation"`
	SensitivePopulation string `json:"sensitivePopulation"`
}

type AirPollutantRef struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}

type AirPollutant struct {
	Code         string           `json:"code"`
	Name         string           `json:"name"`
	FullName     string           `json:"fullName"`
	Concentration Concentration   `json:"concentration"`
}

type Concentration struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

func (c *Client) GetAirQuality(lat, lon float64) (*AirQualityResult, error) {
	latStr := fmt.Sprintf("%.2f", lat)
	lonStr := fmt.Sprintf("%.2f", lon)
	body, err := c.doGet(c.airBaseURL(), fmt.Sprintf("/airquality/v1/current/%s/%s", latStr, lonStr), nil)
	if err != nil {
		return nil, err
	}
	var result AirQualityResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse air quality: %w", err)
	}
	return &result, nil
}
