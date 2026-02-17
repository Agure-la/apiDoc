package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/models"

)

func parseSchemas(doc *openapi3.T) []models.Schema {
	var schemas []models.Schema

	for name, schemaRef := range doc.Components.Schemas {
		schema := schemaRef.Value

		schemas = append(schemas, models.Schema{
			Name:        name,
			Type:        models.SchemaType(schema.Type),
			Description: schema.Description,
			Required:    schema.Required,
		})
	}

	return schemas
}
