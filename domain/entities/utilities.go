package entities

import (
	"encoding/json"
	"net/http"
)

type Answer struct {
	StatusCode int         `json:"satus_code"`
	Message    interface{} `json:"message"`
}

func AccountsAnswer(statusCode int, message interface{}, w http.ResponseWriter) {
	answ := Answer{
		StatusCode: statusCode,
		Message:    message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(answ)
}
