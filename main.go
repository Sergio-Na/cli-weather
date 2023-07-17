package main

import (
	"flag"
	"fmt"
)

func main() {

	locationPtr := flag.String("location", "", "The name of the location for which you would like to check the weather")

	flag.Parse()

	if *locationPtr == "" {
		fmt.Println("An empty location value is invalid")
	} else {
		fmt.Println("The location you chose is:", *locationPtr)
	}

}
