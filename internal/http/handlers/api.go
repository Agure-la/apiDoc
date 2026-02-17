package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/agure-la/api-docs/internal/errors"
	"github.com/agure-la/api-docs/internal/models"
	"github.com/agure-la/api-docs/internal/response"
	"github.com/agure-la/api-docs/internal/spec"
)

type APIHandler struct {
	service *spec.Service
}

func NewAPIHandler(service *spec.Service) *APIHandler {
	return &APIHandler{
		service: service,
	}
}

func (h *APIHandler) ListAPIs(w http.ResponseWriter, r *http.Request) {
	apis := h.service.GetAPIs()
	response.WriteJSONResponse(w, http.StatusOK, apis)
}

func (h *APIHandler) GetAPI(w http.ResponseWriter, r *http.Request) {
	apiName := strings.TrimPrefix(r.URL.Path, "/apis/")
	if apiName == "" {
		response.WriteError(w, http.StatusBadRequest, "API name is required")
		return
	}

	api, err := h.service.GetAPI(apiName)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.Code == errors.ErrNotFound {
			response.WriteError(w, http.StatusNotFound, err.Error())
		} else {
			response.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, api)
}

func (h *APIHandler) ListAPIVersions(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 {
		response.WriteError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	apiName := pathParts[1]
	versions, err := h.service.GetAPIVersions(apiName)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, versions)
}

func (h *APIHandler) GetAPIVersion(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 4 {
		response.WriteError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	apiName := pathParts[1]
	version := pathParts[3]

	apiVersion, err := h.service.GetAPIVersion(apiName, version)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, apiVersion)
}

func (h *APIHandler) CreateAPI(w http.ResponseWriter, r *http.Request) {
	var req models.CreateAPIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Basic validation
	if req.Name == "" || req.Title == "" || req.Version == "" {
		response.WriteError(w, http.StatusBadRequest, "name, title, and version are required")
		return
	}

	doc, err := h.service.CreateAPI(&req)
	if err != nil {
		if apiErr, ok := err.(*errors.APIError); ok && apiErr.Code == errors.ErrConflict {
			response.WriteError(w, http.StatusConflict, err.Error())
		} else {
			response.WriteError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	createResponse := models.CreateAPIResponse{
		ID:      doc.Name,
		Name:    doc.Name,
		Version: req.Version,
		Message: "API created successfully",
	}

	response.WriteJSONResponse(w, http.StatusCreated, createResponse)
}

func (h *APIHandler) UpdateAPI(w http.ResponseWriter, r *http.Request) {
	apiName := strings.TrimPrefix(r.URL.Path, "/apis/")
	if apiName == "" {
		response.WriteError(w, http.StatusBadRequest, "API name is required")
		return
	}

	var req models.UpdateAPIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	doc, err := h.service.UpdateAPI(apiName, &req)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response.WriteJSONResponse(w, http.StatusOK, doc)
}

func (h *APIHandler) DeleteAPI(w http.ResponseWriter, r *http.Request) {
	apiName := strings.TrimPrefix(r.URL.Path, "/apis/")
	if apiName == "" {
		response.WriteError(w, http.StatusBadRequest, "API name is required")
		return
	}

	if err := h.service.DeleteAPI(apiName); err != nil {
		response.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	response := map[string]string{
		"message": "API deleted successfully",
		"name":    apiName,
	}

	response.WriteJSONResponse(w, http.StatusOK, response)
}
