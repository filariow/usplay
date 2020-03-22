package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// readAllTodoHandler ReadAll Todo Handler for REST API
type readAllTodoHandler struct {
	repo Repository
}

//NewReadAllTodoHandler readAlls a new readAll todo handler
func NewReadAllTodoHandler(repo Repository) http.Handler {
	return &readAllTodoHandler{
		repo: repo,
	}
}

// ServeHTTP readAll todo handler
func (h *readAllTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// readAll todo
	todos, err := h.repo.ReadAll()
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:         http.StatusInternalServerError,
			Message:      "Can not readAll todo in database",
			ErrorMessage: err.Error(),
		}
		writeError(w, errMsg)
		return
	}

	// respond to user
	todoData, err := json.Marshal(todos)
	if err != nil {
		log.Printf("Error producing response: %v", err)
		return
	}
	if _, err := w.Write(todoData); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}
