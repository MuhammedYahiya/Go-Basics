package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Declare a global WaitGroup
var Wg sync.WaitGroup

// Function to get the HTTP status code for a given endpoint
func getStatusCode(endpoint string) {
	// Defer the Done() method call to mark the completion of this goroutine
	defer Wg.Done()

	// Make an HTTP GET request to the endpoint
	res, err := http.Get(endpoint)
	if err != nil {
		// Print an error message if the request fails
		fmt.Println("Error occurred for endpoint:", err)
	}

	// Print the status code for the endpoint
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

func main() {
	// List of websites to check
	webSiteList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://go.dev",
	}

	// Loop through the website list
	for _, web := range webSiteList {
		// Start a goroutine to get the status code for each website
		go getStatusCode(web)

		// Increment the WaitGroup counter to indicate the start of a new goroutine
		Wg.Add(1)
	}

	// Wait for all goroutines to finish
	Wg.Wait()
}
