package spec

type Schema struct {
	Name        string           `json:"name"`
	Type        SchemaType       `json:"type"`
	Description string           `json:"description,omitempty"`
	Required    []string         `json:"required,omitempty"`
	Properties  []SchemaProperty `json:"properties,omitempty"`
	Example     any              `json:"example,omitempty"`
}

type SchemaProperty struct {
	Name        string     `json:"name"`
	Type        SchemaType `json:"type"`
	Format      string     `json:"format,omitempty"`
	Description string     `json:"description,omitempty"`
	Required    bool       `json:"required"`
	Example     any        `json:"example,omitempty"`
}

type SchemaType string

const (
	SchemaObject  SchemaType = "object"
	SchemaArray   SchemaType = "array"
	SchemaString  SchemaType = "string"
	SchemaNumber  SchemaType = "number"
	SchemaInteger SchemaType = "integer"
	SchemaBoolean SchemaType = "boolean"
)
