package models

type CreateAPIRequest struct {
	Name        string                 `json:"name" binding:"required"`
	Title       string                 `json:"title" binding:"required"`
	Description string                 `json:"description"`
	Version     string                 `json:"version" binding:"required"`
	Spec        map[string]interface{} `json:"spec" binding:"required"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
}

type UpdateAPIRequest struct {
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Version     string                 `json:"version,omitempty"`
	Spec        map[string]interface{} `json:"spec,omitempty"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
}

type CreateAPIResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Message string `json:"message"`
}
