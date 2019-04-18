package main

import (
	"log"
	"net/http"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/config"
	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/router"
)

func main() {

	config.InitConnectionDB()
	router := router.InitRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}
