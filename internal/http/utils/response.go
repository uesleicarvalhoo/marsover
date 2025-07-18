package utils

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func WriteJSON(w http.ResponseWriter, body []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func WriteError(w http.ResponseWriter, err APIError) {
	payload, jsonErr := json.Marshal(err)
	if jsonErr != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	WriteJSON(w, payload, err.Code)
}

func New(code int, message string, details any) APIError {
	return APIError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

func WriteText(w http.ResponseWriter, body string, statusCode int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(body))
}
