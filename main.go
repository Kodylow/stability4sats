package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kodylow/stability4sats/frontend"
	"github.com/kodylow/stability4sats/handlers"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Load the API key and base URL
	apiKey := os.Getenv("STABILITY_API_KEY")
	if apiKey == "" {
		log.Fatal("STABILITY_API_KEY environment variable is not set")
	}
	baseURL := os.Getenv("API_HOST")
	if baseURL == "" {
		baseURL = "https://api.stability.ai"
	}

	// Pass the API key and base URL to the handlers
	handlers.Init(apiKey, baseURL)

	// Set up the router
	r := mux.NewRouter()

	// Define the endpoints
    r.HandleFunc("/", frontend.IndexHandler)
	r.HandleFunc("/v1/user/account", handlers.UserAccountHandler)
	r.HandleFunc("/v1/user/balance", handlers.UserBalanceHandler)
	r.HandleFunc("/v1/engines/list", handlers.EnginesListHandler)
    r.HandleFunc("/v1/generation/{engine_id}/text-to-image", handlers.TextToImageHandler)

    log.Println("Listening on port 8000")
	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
