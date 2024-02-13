package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{
	"test",
}

var mut sync.Mutex
var Wg sync.WaitGroup

func getStatusCode(endpoint string) {
	defer Wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error occurred for endpoint:", err)
	}
	mut.Lock()                          // Lock the mutex to ensure exclusive access to shared data
	signals = append(signals, endpoint) // Append the endpoint to the signals slice
	mut.Unlock()                        // Unlock the mutex to allow other goroutines to access the shared data

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}

func main() {
	webList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://go.dev",
	}
	for _, web := range webList {
		go getStatusCode(web)
		Wg.Add(1)
	}
	Wg.Wait()
	fmt.Println(signals)
}
