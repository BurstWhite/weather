package handler

import (
	"net/http"
	"strconv"
	"weather-app/backend/internal/service"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	svc *service.WeatherService
}

func NewWeatherHandler(svc *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{svc: svc}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid city id"})
		return
	}

	resp, err := h.svc.GetWeather(id, 0, 0, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
