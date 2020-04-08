package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// updateTodoHandler Update Todo Handler for REST API
type updateTodoHandler struct {
	repo Repository
}

//NewUpdateTodoHandler updates a new update todo handler
func NewUpdateTodoHandler(repo Repository) http.Handler {
	return &updateTodoHandler{
		repo: repo,
	}
}

// ServeHTTP update todo handler
func (h *updateTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	oldTid, err := h.repo.Update(*todo)
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
	todoData, err := json.Marshal(oldTid)
	if err != nil {
		log.Printf("Error producing response: %v", err)
		return
	}
	if _, err := w.Write(todoData); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}

func readTodoBody(r *http.Request) (*Todo, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	var todo Todo
	if err = json.Unmarshal(body, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func writeError(w http.ResponseWriter, errMsg ErrorMessage) {
	jErrMsg, _ := json.Marshal(errMsg)
	w.WriteHeader(http.StatusBadRequest)
	if _, err := w.Write(jErrMsg); err != nil {
		log.Printf("Error writing back error response: %v", err)
	}
}
