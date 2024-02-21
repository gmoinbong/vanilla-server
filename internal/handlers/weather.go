package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiKey = "f9f996577eb403333e3a10667f6b862c"
const units = "metric"

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func GetWeatherData(city string) WeatherData {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=%s",
		city, apiKey, units)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching weather data:", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return WeatherData{}
	}

	var weatherData WeatherData

	err = json.Unmarshal(body, &weatherData)

	if err != nil {
		fmt.Println("error unmarshaling weather data:", err)
		return WeatherData{}
	}
	fmt.Printf("The current temperature in %s is %.2f°C\n", weatherData.Name, weatherData.Main.Temp)
	return weatherData

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
