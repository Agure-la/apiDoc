package spec

type Endpoint struct {
	ID          string        `json:"id"`
	Method      HTTPMethod    `json:"method"`
	Path        string        `json:"path"`
	Summary     string        `json:"summary"`
	Description string        `json:"description,omitempty"`
	Tags        []string      `json:"tags,omitempty"`

	Parameters  []Parameter   `json:"parameters,omitempty"`
	RequestBody *RequestBody `json:"requestBody,omitempty"`
	Responses   []Response    `json:"responses"`

	Deprecated  bool          `json:"deprecated"`
}

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

type Parameter struct {
	Name        string   `json:"name"`
	In          ParamIn  `json:"in"`
	Required    bool     `json:"required"`
	Description string   `json:"description,omitempty"`
	SchemaRef   string   `json:"schemaRef"`
}

type ParamIn string

const (
	ParamPath   ParamIn = "path"
	ParamQuery  ParamIn = "query"
	ParamHeader ParamIn = "header"
)
