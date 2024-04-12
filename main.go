package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Define the URL of the API you want to call
	apiUrl := "https://colosach.onrender.com/api/health-checker"

	// Define a ticker that ticks every 10 minutes
	ticker := time.NewTicker(14 * time.Minute)

	// Call the API immediately upon starting the script
	callAPI(apiUrl)

	// Start a goroutine to periodically call the API
	go func() {
		for range ticker.C {
			callAPI(apiUrl)
		}
	}()

	// Keep the program running indefinitely
	select {}
}

func callAPI(url string) {
	// Make an HTTP GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected response status code:", resp.StatusCode)
		return
	}

	// Print the response body
	// You can add your processing logic here
	fmt.Println("API response:", resp.Status)
}
