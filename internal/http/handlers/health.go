package handlers

import (
	"net/http"

	"github.com/agure-la/api-docs/internal/utils"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ok",
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ready",
	}
	utils.WriteJSONResponse(w, http.StatusOK, response)
}
