package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	log.Println("Started server")
	router := configuredRouter()

	mainRouter := handlers.CombinedLoggingHandler(os.Stdout, router)
	log.Fatalln(http.ListenAndServe(":8080", mainRouter))
}
