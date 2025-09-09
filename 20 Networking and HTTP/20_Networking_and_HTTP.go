package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// startServer starts a simple HTTP server in a goroutine.
func startServer() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
	})

	fmt.Println("Server starting on :8080...")
	// Use a goroutine to run the server so it doesn't block main
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	// Give the server a moment to start up
	time.Sleep(time.Millisecond * 100)
}

// makeClientRequest makes an HTTP GET request to the server.
func makeClientRequest() {
	fmt.Println("Making client request...")
	resp, err := http.Get("http://localhost:8080/hello?name=GoUser")
	if err != nil {
		log.Fatalf("Error making request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	fmt.Printf("Client received: %s\n", body)
}

func main() {
	// Start the server
	startServer()

	// Make a client request after the server has started
	makeClientRequest()

	fmt.Println("Program finished.")
	// In a real application, you might want to keep the main goroutine alive
	// to allow the server to continue running, e.g., by waiting on a channel
	// or a signal.
}
