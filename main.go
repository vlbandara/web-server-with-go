package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// Middleware for logging requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// Handler for the root path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// Handler for a greeting endpoint with query parameters
func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

// JSON response struct
type jsonResponse struct {
	Message string `json:"message"`
}

// Handler for a JSON API endpoint
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{Message: "Hello, JSON World!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Serve static files
func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static"+r.URL.Path)
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Add routes and their handlers
	r.HandleFunc("/", helloHandler).Methods("GET")
	r.HandleFunc("/greet", greetHandler).Methods("GET")
	r.HandleFunc("/api/json", jsonHandler).Methods("GET")

	// Serve static files from the "static" directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Apply logging middleware
	r.Use(loggingMiddleware)

	// Start the server
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server at %s\n", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
