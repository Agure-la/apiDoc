package models

type RequestBody struct {
	Description string `json:"description,omitempty"`
	SchemaRef   string `json:"schemaRef"`
	Required    bool   `json:"required"`
}
