package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/spec"

)

func parseSchemas(doc *openapi3.T) []spec.Schema {
	var schemas []spec.Schema

	for name, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value

		schemas = append(schemas, spec.Schema{
			Name:        name,
			Type:        spec.SchemaType(schema.Type),
			Description: schema.Description,
			Required:    schema.Required,
		})
	}

	return schemas
}
