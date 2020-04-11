package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// deleteTodoHandler Delete Todo Handler for REST API
type deleteTodoHandler struct {
	repo Repository
}

//NewDeleteTodoHandler deletes a new delete todo handler
func NewDeleteTodoHandler(repo Repository) http.Handler {
	return &deleteTodoHandler{
		repo: repo,
	}
}

// ServeHTTP delete todo handler
func (h *deleteTodoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	// delete todo
	ftodo, err := h.repo.Delete(tid)
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:           http.StatusInternalServerError,
			Message:        "Can not delete todo in database",
			ErrorMessage:   err.Error(),
			AdditionalData: map[string]interface{}{"todo-id": tid},
		}
		if IsTodoNotPresentError(err) {
			errMsg.Code = http.StatusNotFound
		}
		writeError(w, errMsg)
		return
	}

	// respond to user
	todoData, err := json.Marshal(ftodo)
	if err != nil {
		log.Printf("Error producing response: %v", err)
		return
	}
	if _, err := w.Write(todoData); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}

func (h *deleteTodoHandler) readPathArgs(r *http.Request) (uuid.UUID, error) {
	params := mux.Vars(r)
	tid, ok := params["tid"]
	if !ok {
		return uuid.UUID{}, fmt.Errorf("Can not find Todo id in request")
	}

	return uuid.Parse(tid)
}
