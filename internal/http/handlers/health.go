package handlers

import (
	"net/http"

	"github.com/agure-la/api-docs/internal/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	healthResponse := map[string]string{
		"status": "ok",
	}
	response.WriteJSONResponse(w, http.StatusOK, healthResponse)
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	readyResponse := map[string]string{
		"status": "ready",
	}
	response.WriteJSONResponse(w, http.StatusOK, readyResponse)
}
