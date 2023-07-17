package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"os"

	"github.com/joho/godotenv"
)

type Location struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	TZID    string  `json:"tz_id"`
}

type Current struct {
	LastUpdatedEpoch int     `json:"last_updated_epoch"`
	LastUpdated      string  `json:"last_updated"`
	TempC            float64 `json:"temp_c"`
	TempF            float64 `json:"temp_f"`
	IsDay            int     `json:"is_day"`
	Condition        struct {
		Text string `json:"text"`
		Icon string `json:"icon"`
		Code int    `json:"code"`
	} `json:"condition"`
	WindMph         float64 `json:"wind_mph"`
	WindKph         float64 `json:"wind_kph"`
	WindDegree      int     `json:"wind_degree"`
	WindDir         string  `json:"wind_dir"`
	PressureMb      float64 `json:"pressure_mb"`
	PressureIn      float64 `json:"pressure_in"`
	PrecipMm        float64 `json:"precip_mm"`
	PrecipIn        float64 `json:"precip_in"`
	Humidity        int     `json:"humidity"`
	Cloud           int     `json:"cloud"`
	FeelsLikeC      float64 `json:"feelslike_c"`
	FeelsLikeF      float64 `json:"feelslike_f"`
	VisibilityKm    float64 `json:"vis_km"`
	VisibilityMiles float64 `json:"vis_miles"`
	UV              float64 `json:"uv"`
	GustMph         float64 `json:"gust_mph"`
	GustKph         float64 `json:"gust_kph"`
}

type WeatherResponse struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func main() {

	locationPtr := flag.String("location", "", "The name of the location for which you would like to check the weather")
	flag.Parse()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	weatherAPIkey := os.Getenv("WEATHER_API_KEY")

	if *locationPtr == "" {
		fmt.Println("An empty location value is invalid")
	} else {

		params := url.Values{}

		params.Add("key", weatherAPIkey)
		params.Add("q", *locationPtr)
		params.Add("aqi", "no")

		resp, err := http.Get("http://api.weatherapi.com/v1/current.json?" + params.Encode())

		if err != nil {
			log.Printf("Request Failed: %s", err)
			return
		}

		defer resp.Body.Close()

		var result WeatherResponse
		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			log.Fatalf("Failed to decode JSON response: %v", err)
		}

		fmt.Println("Name:", result.Location.Name)
		fmt.Println("Region:", result.Location.Region)
		fmt.Println("Country:", result.Location.Country)
		fmt.Println("Latitude:", result.Location.Lat)
		fmt.Println("Longitude:", result.Location.Lon)

		fmt.Println("\nCurrent Weather:")
		fmt.Println("Last Updated:", result.Current.LastUpdated)
		fmt.Println("Temperature (Celsius):", result.Current.TempC)
		fmt.Println("Temperature (Fahrenheit):", result.Current.TempF)
		fmt.Println("Condition:", result.Current.Condition.Text)
		fmt.Println("Wind Speed (mph):", result.Current.WindMph)
		fmt.Println("Wind Speed (kph):", result.Current.WindKph)
	}

}
