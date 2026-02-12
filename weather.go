package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func init() {
	// This prints the moment the package is imported by someone else
	fmt.Printf("[WEATHER PACKAGE] Author: MY-NAME | Loaded at: %s\n", time.Now().Format("15:04:05"))
}

// WeatherResponse matches the structure of the Open-Meteo API
type WeatherResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		Windspeed   float64 `json:"windspeed"`
	} `json:"current_weather"`
}

// GetTemp takes lat/lon and returns the temperature in Celsius
func GetTemp(lat, lon float64) (float64, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, lon)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("network error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, fmt.Errorf("decode error: %w", err)
	}

	return data.CurrentWeather.Temperature, nil
}
