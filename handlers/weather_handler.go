package handlers

import (
	"net/http"
	"weather-check-app/utils"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	location := c.Query("q")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location query parameter is required"})
		return
	}

	temperature, err := utils.GroupedWeatherQuery(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch weather data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"location":    location,
		"temperature": temperature,
	})
}
