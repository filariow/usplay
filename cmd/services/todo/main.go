package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Started server")
	router := configuredRouter()

	mainRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	log.Fatalln(http.ListenAndServe(":8080", mainRouter))
}

func configuredRouter() *mux.Router {
	repo := NewInMemoryRepository()
	router := mux.NewRouter()

	router.NewRoute().Subrouter().
		Methods(http.MethodPost).
		Path("/todo").
		Handler(NewCreateTodoHandler(repo))

	router.NewRoute().Subrouter().
		Methods(http.MethodGet).
		Path("/todo/{tid}").
		Handler(NewReadTodoHandler(repo))

	router.NewRoute().Subrouter().
		Methods(http.MethodGet).
		Path("/todos").
		Handler(NewReadAllTodoHandler(repo))

	router.NewRoute().Subrouter().
		Methods(http.MethodDelete).
		Path("/todo/{tid}").
		Handler(NewDeleteTodoHandler(repo))

	router.NewRoute().Subrouter().
		Methods(http.MethodPut).
		Path("/todo").
		Handler(NewUpdateTodoHandler(repo))

	return router
}
