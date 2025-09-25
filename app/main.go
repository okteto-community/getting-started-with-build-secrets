package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//go:embed name.txt
var nameFileContent string

var loadedName string

// loadEmbeddedName loads the name from the embedded file
func loadEmbeddedName() {
	// Trim whitespace and store the name
	loadedName = strings.TrimSpace(nameFileContent)
	log.Printf("Loaded embedded name: %s", loadedName)
}

// helloHandler handles the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if loadedName == "" {
		http.Error(w, "No name loaded", http.StatusInternalServerError)
		return
	}

	response := fmt.Sprintf("hello %s", loadedName)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

// healthHandler provides a simple health check
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	// Load name from embedded file
	loadEmbeddedName()

	// Set up routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)

	// Start server
	port := ":8080"
	log.Printf("Starting server on port %s", port)
	log.Printf("Endpoints available:")
	log.Printf("  GET /hello - Returns 'hello {name}'")
	log.Printf("  GET /health - Health check")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
