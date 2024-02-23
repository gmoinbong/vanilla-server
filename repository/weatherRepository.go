package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"vanilla-server/internal/handlers"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const (
	weatherCacheKey       = "weather:%s"
	forecastWeatherKeyFmt = "forecast:%s"
	cacheExpiration       = 5 * time.Minute 
)

type WeatherRepository struct {
	redisClient *redis.Client
}

func NewWeatherRepository(redisClient *redis.Client) *WeatherRepository {
	return &WeatherRepository{redisClient: redisClient}
}

func (r *WeatherRepository) GetWeatherByCity(city string) (*handlers.WeatherData, error) {

	weatherKey := fmt.Sprintf(weatherCacheKey, city)
	cachedWeather, err := r.redisClient.Get(ctx, weatherKey).Result()
	if err == nil {
		var weatherData handlers.WeatherData
		if err := json.Unmarshal([]byte(cachedWeather), &weatherData); err != nil {
			return nil, err
		}
		return &weatherData, nil
	} else if err != redis.Nil {
		return nil, err 
	}

	weatherData, err := handlers.GetCurrentWeatherData()
	if err != nil {
		return nil, err 
	}

	weatherJSON, err := json.Marshal(weatherData)
	if err != nil {
		return nil, err 
	}
	if err := r.redisClient.Set(ctx, weatherKey, weatherJSON, cacheExpiration).Err(); err != nil {
		return nil, err 
	}

	return &weatherData, nil
}

func (r *WeatherRepository) GetForecastWeatherByCity(city string) (*handlers.ForecastWeatherFiveDays, error) {

	forecastWeatherKey := fmt.Sprintf(forecastWeatherKeyFmt, city)
	cachedForecastWeather, err := r.redisClient.Get(ctx, forecastWeatherKey).Result()
	if err == nil {
		var forecastWeatherData handlers.ForecastWeatherFiveDays
		if err := json.Unmarshal([]byte(cachedForecastWeather), &forecastWeatherData); err != nil {
			return nil, err
		}
		return &forecastWeatherData, nil
	} else if err != redis.Nil {
		return nil, err 
	}

	forecastWeatherData, err := handlers.GetForecastWeatherFiveDays()
	if err != nil {
		return nil, err 
	}

	forecastWeatherJSON, err := json.Marshal(forecastWeatherData)
	if err != nil {
		return nil, err
	}
	if err := r.redisClient.Set(ctx, forecastWeatherKey, forecastWeatherJSON, cacheExpiration).Err(); err != nil {
		return nil, err 
	}

	return &forecastWeatherData, nil
}
// Время жизни кэша: 5 минут