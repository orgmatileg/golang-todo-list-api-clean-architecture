package router

import (
	"fmt"

	"github.com/orgmatileg/golang-todo-list-api-clean-architecture/config"

	// Todos

	"net/http"

	"github.com/gorilla/mux"
	m "github.com/orgmatileg/golang-todo-list-api-clean-architecture/middleware"

	hTodos "github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/delivery/http"
	_todosRepo "github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/repository"
	_todosUcase "github.com/orgmatileg/golang-todo-list-api-clean-architecture/module/todos/usecase"
)

// InitRouter endpoint
func InitRouter() *mux.Router {

	r := mux.NewRouter()

	// Check API
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	}).Methods("GET")

	// Endpoint for testing function or such a thing
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test!")
	}).Methods("POST")

	// Init versioning API
	rv1 := r.PathPrefix("/v1").Subrouter()

	// Middleware
	rv1.Use(m.CORS)

	// Get DB Conn
	dbConn := config.GetPostgresDB()

	// Todos
	todoRepo := _todosRepo.NewTodosRepositoryPostgres(dbConn)
	todoUcase := _todosUcase.NewTodosUsecase(todoRepo)
	hTodos.NewTodosHttpHandler(rv1, todoUcase)

	return r
}
