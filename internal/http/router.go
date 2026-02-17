package http

import (
	"net/http"
	"strings"

	"github.com/agure-la/api-docs/internal/spec"
)

type Router struct {
	apiHandler    *APIHandler
	healthHandler *HealthHandler
}

func NewRouter(service *spec.Service) *Router {
	return &Router{
		apiHandler:    NewAPIHandler(service),
		healthHandler: NewHealthHandler(),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Health endpoints
	mux.HandleFunc("/health", r.healthHandler.Health)
	mux.HandleFunc("/ready", r.healthHandler.Ready)

	// API endpoints
	mux.HandleFunc("/apis", r.handleAPIs)
	mux.HandleFunc("/apis/", r.handleAPIs)

	return mux
}

func (r *Router) handleAPIs(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		path := req.URL.Path
		
		// /apis - list all APIs
		if path == "/apis" {
			r.apiHandler.ListAPIs(w, req)
			return
		}
		
		// /apis/{api} - get specific API
		if len(path) > len("/apis/") && !strings.Contains(path[len("/apis/"):], "/") {
			r.apiHandler.GetAPI(w, req)
			return
		}
		
		// /apis/{api}/versions - list versions for an API
		if strings.HasSuffix(path, "/versions") {
			r.apiHandler.ListAPIVersions(w, req)
			return
		}
		
		// /apis/{api}/versions/{version} - get specific version
		pathParts := strings.Split(strings.Trim(path, "/"), "/")
		if len(pathParts) == 4 && pathParts[2] == "versions" {
			r.apiHandler.GetAPIVersion(w, req)
			return
		}
		
		WriteError(w, http.StatusNotFound, "Not found")
	default:
		WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
