package main

import (
	"log"

	"br.com.charlesrodrigo/ms-person/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.StartServerApi()

}
