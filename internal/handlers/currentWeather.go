package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vanilla-server/utils"
)

const apiKey = "f9f996577eb403333e3a10667f6b862c"
const units = "metric"

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func GetCurrentWeatherData() (WeatherData, error) {
	city, err := utils.CityInputReader()

	if err != nil {
		return WeatherData{}, fmt.Errorf("error reading city input :%v", err)
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s",
		city, apiKey, units)

	res, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error fetching weather data: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error reading response body: %v", err)
	}

	var weatherData WeatherData

	err = json.Unmarshal(body, &weatherData)

	if err != nil {
		return WeatherData{}, fmt.Errorf("error unmarshaling weather data: %v", err)
	}
	fmt.Printf("The current temperature in %s is %.2f°C\n", weatherData.Name, weatherData.Main.Temp)
	return weatherData, nil

}
func CurrentWeatherDataHandler(w http.ResponseWriter, r *http.Request) {
	weatherData, err := GetCurrentWeatherData()
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting weather data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherData)

}

func DisplayWeatherData(db *sql.DB) {
	rows, err := db.Query("SELECT city, temperature FROM weather")
	if err != nil {
		fmt.Println("Error querying from weather data:", err)
	}
	defer rows.Close()

	for rows.Next() {
		var city string
		var temperature float64

		err := rows.Scan(&city, &temperature)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("%s: %.2f°C\n", city, temperature)
	}
}
