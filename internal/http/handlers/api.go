package handlers

import (
	"net/http"
	"strings"

	"github.com/agure-la/api-docs/internal/spec"
	"github.com/agure-la/api-docs/internal/utils"
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
	utils.WriteJSONResponse(w, http.StatusOK, apis)
}

func (h *APIHandler) GetAPI(w http.ResponseWriter, r *http.Request) {
	apiName := strings.TrimPrefix(r.URL.Path, "/apis/")
	if apiName == "" {
		utils.WriteError(w, http.StatusBadRequest, "API name is required")
		return
	}

	api, err := h.service.GetAPI(apiName)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, api)
}

func (h *APIHandler) ListAPIVersions(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 3 {
		utils.WriteError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	apiName := pathParts[1]
	versions, err := h.service.GetAPIVersions(apiName)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, versions)
}

func (h *APIHandler) GetAPIVersion(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) < 4 {
		utils.WriteError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	apiName := pathParts[1]
	version := pathParts[3]

	apiVersion, err := h.service.GetAPIVersion(apiName, version)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, apiVersion)
}
