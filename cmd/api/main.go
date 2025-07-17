package main

import (
	"fmt"
	"log"
	

	"runeDiscordBot/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}

	fmt.Println("Running in API-only mode...")
	api.SetupAPI()
}