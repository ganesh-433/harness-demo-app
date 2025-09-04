// main.go (for Canary Version)
package main

import (
	"fmt"
	"net/http"
	"os"
)

var version = "2.0.0"

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, Canary! I'm version %s on host %s\n", version, hostname)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}