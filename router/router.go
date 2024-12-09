package router

import (
	"github.com/gin-gonic/gin"
	"weather-check-app/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/weather", handlers.GetWeather)
	return r
}
