package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
	"vanilla-server/utils"
)

type ForecastWeatherFiveDays struct {
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

func GetForecastWeatherFiveDays() (string, error) {
	city, err := utils.CityInputReader()

	if err != nil {
		return "", fmt.Errorf("error reading city input :%v", err)
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&dt=6&appid=%s&units=metric", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching forecast of five days :%v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body :%v", err)
	}

	var forecastWeatherData ForecastWeatherFiveDays

	err = json.Unmarshal(body, &forecastWeatherData)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling weather data :%v ", err)
	}

	return formatForecast(forecastWeatherData), nil
}
func parseForecastTime(timeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", timeStr)
}

func formatForecast(data ForecastWeatherFiveDays) string {
	uniqueDays := make(map[string]bool)
	var formattedForecast string
	for _, forecastData := range data.List {
		t, err := parseForecastTime(forecastData.DtTxt)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			continue
		}
		if t.Hour() == 12 && !uniqueDays[t.Format("2006-01-02")] {
			uniqueDays[t.Format("2006-01-02")] = true

			formattedForecast += fmt.Sprintf("%s, Temperature: %.2f°C, Humidity: %.2f\n", t.Format(time.UnixDate), forecastData.Main.Temp, forecastData.Main.Humidity)
		}
	}
	return formattedForecast

}

func ForecastWeatherHandler(w http.ResponseWriter, r *http.Request) {
	forecastData, err := GetForecastWeatherFiveDays()

	if err != nil {
		http.Error(w, fmt.Sprintf("error getting forecast :%v", err), http.StatusInternalServerError)
		return
	}
	if forecastData == "" {
		http.Error(w, "forecast data empty", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(map[string]string{"forecast": forecastData})
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshaling forecast data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)

	fmt.Println("Forecast data: \n", forecastData)
}
