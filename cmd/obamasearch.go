package main

import (
	"github.com/joho/godotenv"
	"github.com/matthewrudy/obamasearch"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	obamasearch.Search("obama")
}
