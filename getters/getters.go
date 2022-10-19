package getters

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type WeatherGetter interface {
	IsRaining(float64, float64) bool
	GetTemperature(float64, float64) float64
	GetWindSpeed(float64, float64) float64
}

type ActualWeather struct{}

func (ActualWeather) IsRaining(lat float64, lon float64) bool {
	URL := buildBaseURL(lat, lon)
	URL += "&current_weather=true"

	jsonMap := returnMapFromURL(URL)

	weatherCode := int(jsonMap["current_weather"].(map[string]interface{})["weathercode"].(float64))

	if (weatherCode/10 == 5) || (weatherCode/10 == 6) || (weatherCode/10 == 8) {
		return true
	}

	return false

}

func (ActualWeather) GetTemperature(lat float64, lon float64) float64 {
	URL := buildBaseURL(lat, lon)
	URL += "&current_weather=true"

	jsonMap := returnMapFromURL(URL)

	return jsonMap["current_weather"].(map[string]interface{})["temperature"].(float64)
}

func (ActualWeather) GetWindSpeed(lat float64, lon float64) float64 {
	URL := buildBaseURL(lat, lon)
	URL += "&current_weather=true"

	jsonMap := returnMapFromURL(URL)

	return jsonMap["current_weather"].(map[string]interface{})["windspeed"].(float64)
}

type RandomWeather struct{}

func (RandomWeather) IsRaining(lat float64, lon float64) bool {
	return rand.Intn(2) == 1
}

func (RandomWeather) GetTemperature(lat float64, lon float64) float64 {
	return (rand.Float64() * 150.0) - 90.0
}

func (RandomWeather) GetWindSpeed(lat float64, lon float64) float64 {
	return rand.Float64() * 408.0
}

func buildBaseURL(lat float64, lon float64) string {
	URL := "https://api.open-meteo.com/v1/forecast"
	URL += "?latitude="
	URL += strconv.FormatFloat(lat, 'f', 4, 64)
	URL += "&longitude="
	URL += strconv.FormatFloat(lon, 'f', 4, 64)

	return URL
}

func returnMapFromURL(URL string) map[string]interface{} {
	response, _ := http.Get(URL)

	body, _ := io.ReadAll(response.Body)

	var jsonMap map[string]interface{}
	json.Unmarshal(body, &jsonMap)

	return jsonMap

}
