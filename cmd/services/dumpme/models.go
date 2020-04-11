package main

// ErrorMessage error message for communication
type ErrorMessage struct {
	Message        string                 `json:"message"`
	Code           int                    `json:"code"`
	ErrorMessage   string                 `json:"errorMessage"`
	AdditionalData map[string]interface{} `json:"additionalData"`
}
