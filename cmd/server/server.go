package server

import (
	"database/sql"
	"log"
	"net/http"
	"vanilla-server/internal/config"
	"vanilla-server/internal/handlers"
	"vanilla-server/internal/storage"
)

func RunServer() {
	cfg := config.MustLoadConfig()
	db, err := storage.InitDB(cfg)
	if err != nil {
		log.Fatalf("Eror initializing DB: %v", err)
	}
	defer db.Close()

	weatherData := handlers.GetWeatherData("Kyiv")
	if err := insertWeatherData(db, weatherData); err != nil {
		log.Fatalf("Error inserting weather data: %v", err)
	}

	// handlers.DisplayWeatherData(db)
	handlers.GetForecastFiveDays("Donetsk")

	if err = http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func insertWeatherData(db *sql.DB, weatherData handlers.WeatherData) error {
	_, err := db.Exec("INSERT INTO weather (city, temperature) VALUES ($1, $2)", weatherData.Name, weatherData.Main.Temp)
	return err
}
