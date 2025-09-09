# Go Networking and HTTP

Go's standard library provides robust and easy-to-use packages for networking, especially for building HTTP servers and clients. This tutorial demonstrates how to create a simple HTTP server and how to make a client request to it.

## HTTP Server (`net/http`)

The `net/http` package allows you to create powerful web servers with minimal code. Key components include:
-   `http.HandleFunc`: Registers a function to handle requests for a given path.
-   `http.ListenAndServe`: Starts an HTTP server that listens on a specified address.
-   `http.ResponseWriter`: Interface used by an HTTP handler to construct an HTTP response.
-   `*http.Request`: Represents an incoming HTTP request received by a server.

## HTTP Client (`net/http`)

Making HTTP requests from a Go application is also straightforward using the `net/http` package. The `http.Get` function is a convenient way to perform a GET request.

## Example: Simple Server and Client

The `20_Networking_and_HTTP.go` file contains an example where:
1.  A simple HTTP server is started in a goroutine, listening on port 8080 and serving the `/hello` endpoint.
2.  An HTTP client then makes a GET request to this server.

```go
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
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	time.Sleep(time.Millisecond * 100) // Give server time to start
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
	startServer()
	makeClientRequest()
	fmt.Println("Program finished.")
}
```

## Running the Example

To run this example, navigate to the `20 Networking and HTTP` directory in your terminal and execute:

```bash
go run 20_Networking_and_HTTP.go
```

You will see output indicating the server starting, the client making a request, and the response received from the server.
