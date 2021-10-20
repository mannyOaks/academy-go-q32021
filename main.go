package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mannyOaks/academy-go-q32021/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	app.RunApp()
}
