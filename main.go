package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/safianalicb/weather/getters"
)

func main() {

	var weather getters.WeatherGetter
	weather = getters.ActualWeather{}

	var lat, lon float64

	fmt.Printf("Latitude: ")
	_, err := fmt.Scanf("%f", &lat)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Longitude: ")
	_, err = fmt.Scanf("%f", &lon)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Println("Not random:")
	fmt.Printf("\nRaining? %v\n", weather.IsRaining(lat, lon))
	fmt.Printf("The temperature is: %.1f°\n", weather.GetTemperature(lat, lon))
	fmt.Printf("The wind speed is %.1fkm/h\n\n", weather.GetWindSpeed(lat, lon))

	rand.Seed(time.Now().UnixNano())
	weather = getters.RandomWeather{}

	fmt.Println("Random:")
	fmt.Printf("\nRaining? %v\n", weather.IsRaining(lat, lon))
	fmt.Printf("The temperature is: %.1f°\n", weather.GetTemperature(lat, lon))
	fmt.Printf("The wind speed is %.1fkm/h\n", weather.GetWindSpeed(lat, lon))

}
