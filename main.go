package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/FayxChance/tha-backend-tezos/internal/app"
)

func main() {
	port := flag.String("port", "8080", "Port to run the web server on")
	flag.Parse()
	// Start the server on the provided port or default to 8080
	envPort, ok := os.LookupEnv("PORT")
	if ok {
		port = &envPort
	}

	var app app.App
	err := app.SetupApp()
	if err != nil {
		panic(err)
	}

	log.Printf("Server running on port: %s", *port)
	err = app.Router.Router.Run(fmt.Sprintf(":%s", *port)) // Run on the specified port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
