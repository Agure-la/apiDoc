package handlers

import (
	"net/http"

	"github.com/agure-la/api-docs/internal/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ok",
	}
	http.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ready",
	}
	http.WriteJSONResponse(w, http.StatusOK, response)
}
