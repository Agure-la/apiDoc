package http

import (
	"net/http"
	"strings"

	"github.com/agure-la/api-docs/internal/http/handlers"
	"github.com/agure-la/api-docs/internal/response"
	"github.com/agure-la/api-docs/internal/spec"
)

type Router struct {
	apiHandler    *handlers.APIHandler
	healthHandler *handlers.HealthHandler
}

func NewRouter(service *spec.Service) *Router {
	return &Router{
		apiHandler:    handlers.NewAPIHandler(service),
		healthHandler: handlers.NewHealthHandler(),
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
		
		response.WriteError(w, http.StatusNotFound, "Not found")
		
	case http.MethodPost:
		// /apis - create new API
		if req.URL.Path == "/apis" {
			r.apiHandler.CreateAPI(w, req)
			return
		}
		response.WriteError(w, http.StatusNotFound, "Not found")
		
	case http.MethodPut:
		// /apis/{api} - update existing API
		if len(req.URL.Path) > len("/apis/") && !strings.Contains(req.URL.Path[len("/apis/"):], "/") {
			r.apiHandler.UpdateAPI(w, req)
			return
		}
		response.WriteError(w, http.StatusNotFound, "Not found")
		
	case http.MethodDelete:
		// /apis/{api} - delete API
		if len(req.URL.Path) > len("/apis/") && !strings.Contains(req.URL.Path[len("/apis/"):], "/") {
			r.apiHandler.DeleteAPI(w, req)
			return
		}
		response.WriteError(w, http.StatusNotFound, "Not found")
		
	default:
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
