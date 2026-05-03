package handler

import (
	"net/http"
	"strconv"
	"weather-app/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type CityHandler struct{}

func NewCityHandler() *CityHandler {
	return &CityHandler{}
}

func (h *CityHandler) ListCities(c *gin.Context) {
	cities, err := service.ListCities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (h *CityHandler) AddCity(c *gin.Context) {
	var req struct {
		Name       string  `json:"name"`
		LocationID string  `json:"locationId"`
		Adm1       string  `json:"adm1"`
		Adm2       string  `json:"adm2"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cities, _ := service.ListCities()
	for _, city := range cities {
		if city.LocationID == req.LocationID {
			c.JSON(http.StatusOK, city)
			return
		}
	}
	city, err := service.AddCity(req.Name, req.LocationID, req.Adm1, req.Adm2, req.Lat, req.Lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, city)
}

func (h *CityHandler) DeleteCity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := service.DeleteCity(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
