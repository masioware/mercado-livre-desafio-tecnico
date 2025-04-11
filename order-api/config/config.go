package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func GetDistributionCenterURL() string {
	url := os.Getenv("DISTRIBUTION_CENTER_URL")
	if url == "" {
		url = "http://localhost:8001/distribuitioncenters" // fallback
	}
	return url
}
