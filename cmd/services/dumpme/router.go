package main

import (
	"github.com/gorilla/mux"
)

func configuredRouter() *mux.Router {
	router := mux.NewRouter()

	router.NewRoute().Subrouter().
		Path("/dumpme").
		HandlerFunc(DumpRequest)

	return router
}
