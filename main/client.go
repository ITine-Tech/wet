package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func getWeather(location string) (Weather, error) {
	//Create new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	// Check Redis cache for weather data
	ctx := context.Background()
	val, err := client.Get(ctx, location).Result()
	if err == nil {
		var weather Weather
		err = json.Unmarshal([]byte(val), &weather)
		if err == nil {
			return weather, nil
		}
	}
	// Fetch weather data from API if cache is empty or expired
	response, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=" + APIkey + "&q=" + location + "&days=" + days + "&aqi=no&alerts=no")

	if err != nil {
		return Weather{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Weather{}, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return Weather{}, err
	}

	// Set Redis cache with weather data
	err = client.Set(ctx, location, body, cacheExpiration).Err()
	if err != nil {
		log.Printf("Error setting cache: %v\n", err)
	}

	return weather, nil
}
