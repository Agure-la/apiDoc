package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data any `json:"data,omitempty"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		return json.NewEncoder(w).Encode(data)
	}

	return nil
}

func WriteJSONResponse(w http.ResponseWriter, status int, data any) error {
	return WriteJSON(w, status, Response{Data: data})
}

func WriteError(w http.ResponseWriter, status int, message string) error {
	return WriteJSON(w, status, ErrorResponse{Error: message})
}
