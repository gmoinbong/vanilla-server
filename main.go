package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

const apiKey = "186bd86d8ce1b477fbb716010c6199a2"

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Name string `json:"name"`
}

func main() {
	connectionString := fmt.Sprintln("host=localhost port=5432 user=admin password=admin dbname=postgres_server sslmode=disable")

	db, err := sql.Open("postgres", connectionString)
	CheckError("Error opening db:", err)
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS weather (id SERIAL PRIMARY KEY, city TEXT, temperature REAL )")
	CheckError("Error creating table:", err)

	// err = db.Ping()
	// CheckError("ping error", err)
	weatherData := getWeatherData("Kyiv")

	_, err = db.Exec("INSERT INTO weather (city, temperature) VALUES ($1, $2)", weatherData.Name, weatherData.Main.Temp)
	CheckError("Error inserting weather data:", err)
	displayWeatherData(db)
}

func getWeatherData(city string) WeatherData {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)
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

func displayWeatherData(db *sql.DB) {
	rows, err := db.Query("SELECT city, temperature FROM weather")
	CheckError("Error querying from weather data:", err)
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

func CheckError(description string, err error) {
	if err != nil {
		panic(err)
	}
}
