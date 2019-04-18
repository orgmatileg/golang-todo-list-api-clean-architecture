package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"log"
	"net/http"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/config"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/router"
)

func init() {

	if env := os.Getenv("GO_ENV"); env != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	config.InitConnectionDB()

}

func main() {

	router := router.InitRouter()

	log.Println("Server running on Port 8081")

	err := http.ListenAndServe(":8081", router)

	if err != nil {
		log.Println("Oops! Something went wrong on your server!")
		log.Fatal(err)
	}

}
