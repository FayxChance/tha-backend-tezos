package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/FayxChance/tha-backend-tezos/internal/app"
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
	err := app.SetupApp()
	if err != nil {
		log.Fatalf("Failed to set up the app: %v", err)
	}

	// Initialize the delegations service with the database
	delegationsService := app.Router.DelegationCtrl.DelegationsSvc

	// Define the API URL to fetch delegations from
	apiURL := "https://api.tzkt.io/v1/operations/delegations"

	lastID, err := delegationsService.LastDelegationTzktID()
	if err != nil {
		panic(err)
	}

	if lastID != 0 {
		err = delegationsService.FetchFirst1000Delegations(apiURL)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Data existing, doesn't fetch first 1k data.\n")
	}

	go delegationsService.StartFetch(apiURL)

	log.Printf("Server running on port: %s", *port)
	err = app.Router.Router.Run(fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
