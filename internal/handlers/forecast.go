package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Forecast5DaysWeatherData struct {
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	List []struct {
		DtTxt string `json:"dt_txt"`
		Main  struct {
			Temp     float64 `json:"temp"`
			Humidity float64 `json:"humidity"`
		} `json:"main"`
	} `json:"list"`
}

func GetForecastFiveDays(city string) string {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&dt=6&appid=%s&units=metric", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching forecast of five days:", err)
		return ""
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		return ""
	}

	var forecastWeatherData Forecast5DaysWeatherData

	err = json.Unmarshal(body, &forecastWeatherData)
	if err != nil {
		fmt.Println("error unmarshaling weather data:", err)
		return ""
	}

	fmt.Println(formatForecast(forecastWeatherData))

	return formatForecast(forecastWeatherData)
}
func parseForecastTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeStr)
}

func formatForecast(data Forecast5DaysWeatherData) string {
	uniqueDays := make(map[string]bool)
	var formattedForecast string
	for _, data := range data.List {
		t, err := parseForecastTime(data.DtTxt)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		if t.Hour() == 12 && !uniqueDays[t.Format("2006-01-02")] {
			uniqueDays[t.Format("2006-01-02")] = true

			formattedForecast += fmt.Sprintf("%s, Temperature: %.2f°C, Humidity: %.2f\n", t.Format(time.UnixDate), data.Main.Temp, data.Main.Humidity)
		}
	}
	return formattedForecast

}
