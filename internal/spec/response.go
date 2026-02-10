package spec

type Response struct {
	StatusCode  int    `json:"statusCode"`
	Description string `json:"description,omitempty"`
	SchemaRef   string `json:"schemaRef,omitempty"`
}
