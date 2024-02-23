package router

import (
	"net/http"
	"vanilla-server/internal/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/current-weather", handlers.CurrentWeatherDataHandler)
	http.HandleFunc("/forecast", handlers.ForecastWeatherHandler)
}
