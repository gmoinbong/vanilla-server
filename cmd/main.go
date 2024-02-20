package main

import (
	"net/http"
	"vanilla-server/internal/config"
	"vanilla-server/internal/handlers"
	"vanilla-server/internal/storage"
	"vanilla-server/utils"
	"vanilla-server/utils/lockutil"

	_ "github.com/lib/pq"
)

func main() {
	lockutil.RunWithLock(func() {
		cfg := config.MustLoadConfig()
		db, err := storage.InitDB(cfg)
		utils.CheckError("Eror initializing DB: %v", err)

		defer db.Close()

		weatherData := handlers.GetWeatherData("Kyiv")

		_, err = db.Exec("INSERT INTO weather (city, temperature) VALUES ($1, $2)", weatherData.Name, weatherData.Main.Temp)
		utils.CheckError("Error inserting weather data:", err)

		handlers.DisplayWeatherData(db)

		err = http.ListenAndServe(":8082", nil)
		utils.CheckError("Error starting server:", err)
	})
}
