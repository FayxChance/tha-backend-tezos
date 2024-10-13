package main

import (
	"flag" // Used for command-line argument parsing
	"fmt"  // For formatted I/O operations
	"log"  // Logging support
	"os"   // OS-specific functionality (like environment variables)

	"github.com/FayxChance/tha-backend-tezos/internal/app" // Internal app package
)

func main() {
	// Parse the port from command-line arguments, default to 8080
	port := flag.String("port", "8080", "Port to run the web server on")
	flag.Parse()

	// Check if a port is provided by the environment, override if present
	envPort, ok := os.LookupEnv("PORT")
	if ok {
		port = &envPort
	}

	// Initialize the application
	var app app.App
	err := app.SetupApp() // Sets up the app (likely includes database, routers, etc.)
	if err != nil {
		log.Fatalf("Failed to set up the app: %v", err)
	}

	// Initialize the delegations service with the database
	delegationsService := app.Router.DelegationCtrl.DelegationsSvc

	// Define the API URL to fetch delegations from
	apiURL := "https://api.tzkt.io/v1/operations/delegations"

	// Fetch the last delegation ID to avoid fetching duplicate data
	lastID, err := delegationsService.LastDelegationTzktID()
	if err != nil {
		panic(err)
	}

	// Fetch initial data if the last ID is not zero
	if lastID != 0 {
		err = delegationsService.FetchFirst1000Delegations(apiURL)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Data existing, doesn't fetch first 1k data.\n")
	}

	// Start continuous polling for new delegations in a background goroutine
	go delegationsService.StartFetch(apiURL)

	// Start the web server and listen on the configured port
	log.Printf("Server running on port: %s", *port)
	err = app.Router.Router.Run(fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
