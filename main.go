package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	weatherAPIkey := os.Getenv("WEATHER_API_KEY")

	fmt.Println(weatherAPIkey)

	locationPtr := flag.String("location", "", "The name of the location for which you would like to check the weather")

	flag.Parse()

	if *locationPtr == "" {
		fmt.Println("An empty location value is invalid")
	} else {
		fmt.Println("The location you chose is:", *locationPtr)
	}

}
