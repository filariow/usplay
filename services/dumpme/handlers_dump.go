package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// DumpRequest dumps the given request
func DumpRequest(w http.ResponseWriter, r *http.Request) {
	// read request body
	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Println(err)
		errMsg := ErrorMessage{
			Code:         http.StatusInternalServerError,
			Message:      "Internal Server Error",
			ErrorMessage: err.Error(),
		}
		writeError(w, &errMsg)
		return
	}

	// producing response
	resp := &DumpMessageResponse{
		Dump: string(requestDump),
	}
	respJSON, err := json.Marshal(resp)
	if err != nil {
		msg := fmt.Sprintf("Error costructing the json response with for the dump: %s", resp.Dump)
		errMsg := ErrorMessage{
			Code:         http.StatusInternalServerError,
			Message:      msg,
			ErrorMessage: err.Error(),
		}
		writeError(w, &errMsg)
		return
	}

	// sending response
	if _, err := w.Write(respJSON); err != nil {
		log.Printf("Error responding to request: %v", err)
	}
}

func writeError(w http.ResponseWriter, errMsg *ErrorMessage) {
	if payload, err := json.Marshal(errMsg); err != nil {
		log.Printf(
			"Error producing json for errMsg. Code: %v, Message: %s, Error Message: %s",
			errMsg.Code, errMsg.Message, errMsg.ErrorMessage)
	} else {
		w.Write(payload)
	}

	w.WriteHeader(errMsg.Code)
}

// DumpMessageResponse response with the dump of the request
type DumpMessageResponse struct {
	Dump string `json:"dump"`
}
