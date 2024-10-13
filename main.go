package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Delegation struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Level     string    `json:"level"`
}

type Data struct {
	Data []Delegation `json:"data"`
}

func main() {
	port := flag.String("port", "8080", "Port to run the web server on")
	flag.Parse()

	// Create a new Gin router instance
	router := gin.Default()

	// Define a simple GET route
	router.GET("/xtz/delegations", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	// Start the server on the provided port or default to 8080
	envPort, ok := os.LookupEnv("PORT")
	if ok {
		port = &envPort
	}

	log.Printf("Server running on port: %s", *port)

	err := router.Run(fmt.Sprintf(":%s", *port)) // Run on the specified port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
