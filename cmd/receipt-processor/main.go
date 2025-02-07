package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"receipt-processor-challenge/pkg/config"
	"receipt-processor-challenge/pkg/routes"
)

// Function that's the entry point to the application
func main() {
	config.Init()
	router := routes.InitializeRouter()

	port, exists := viper.Get("PORT").(string)
	if !exists {
		port = "8080"
	}

	fmt.Println("Starting Server On Port", port)
	portIsAvailable := isPortAvailable(":" + port)
	if portIsAvailable {
		if err := http.ListenAndServe(":"+port, router); err != nil {
			log.Fatal("Error starting server:", err)
		}
	} else {
		log.Fatalf("Error starting server: port %s is not available", port)
	}
}

// Function to check if port is available and release it if possible
func isPortAvailable(port string) bool {
	portConnection, err := net.Listen("tcp", port)
	if err != nil {
		return false
	}

	err = portConnection.Close()
	if err != nil {
		return false
	}
	return true
}
