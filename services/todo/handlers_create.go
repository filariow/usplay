package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// createTodoHandler Create Todo Handler for REST API
type createTodoHandler struct {
	repo Repository
}

//NewCreateTodoHandler creates a new create todo handler
func NewCreateTodoHandler(repo Repository) http.Handler {
	return &createTodoHandler{
		repo: repo,
	}
}

// ServeHTTP create todo handler
func (h *createTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request input
	todo, err := readTodoBody(r)
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:         http.StatusBadRequest,
			Message:      "Empty or invalid request body",
			ErrorMessage: err.Error(),
		}
		writeError(w, errMsg)
		return
	}

	// store data
	cTodo, err := h.repo.Create(*todo)
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:           http.StatusInternalServerError,
			Message:        "Can not store todo in database",
			ErrorMessage:   err.Error(),
			AdditionalData: map[string]interface{}{"todo": todo},
		}
		writeError(w, errMsg)
		return
	}

	// respond to user
	todoData, err := json.Marshal(cTodo)
	if err != nil {
		log.Printf("Error producing response: %v", err)
		return
	}
	if _, err := w.Write(todoData); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}
