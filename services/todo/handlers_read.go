package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// readTodoHandler Read Todo Handler for REST API
type readTodoHandler struct {
	repo Repository
}

//NewReadTodoHandler reads a new read todo handler
func NewReadTodoHandler(repo Repository) http.Handler {
	return &readTodoHandler{
		repo: repo,
	}
}

// ServeHTTP read todo handler
func (h *readTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// read request input
	tid, err := h.readPathArgs(r)
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
	log.Printf("Parsed id %v", tid)

	// read todo
	todo, err := h.repo.Read(tid)
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:           http.StatusInternalServerError,
			Message:        "Can not read todo in database",
			ErrorMessage:   err.Error(),
			AdditionalData: map[string]interface{}{"todo-id": tid},
		}
		writeError(w, errMsg)
		return
	}

	// respond to user
	if todo == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(tid.String()))
		return
	}

	todoData, err := json.Marshal(todo)
	if err != nil {
		log.Printf("Error producing response: %v", err)
		return
	}
	if _, err := w.Write(todoData); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}

func (h *readTodoHandler) readPathArgs(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	tid, ok := params["tid"]
	if !ok {
		return uuid.UUID{}, fmt.Errorf("Can not find Todo id in request")
	}

	return uuid.Parse(tid)
}
