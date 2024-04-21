package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shamilsdq/ideaboard-go-svelte/server"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Define Server
	server := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: server.NewServer(),
	}

	// Start listening
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
