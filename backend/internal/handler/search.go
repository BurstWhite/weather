package handler

import (
	"net/http"
	"weather-app/backend/internal/qweather"

	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	q *qweather.Client
}

func NewSearchHandler(q *qweather.Client) *SearchHandler {
	return &SearchHandler{q: q}
}

func (h *SearchHandler) SearchCity(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param q required"})
		return
	}
	results, err := h.q.SearchCity(q)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
