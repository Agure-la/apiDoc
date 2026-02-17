package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/models"

)

func parseAuth(doc *openapi3.T) []models.AuthScheme {
	var auth []models.AuthScheme

	for name, scheme := range doc.Components.SecuritySchemes {
		auth = append(auth, models.AuthScheme{
			Type:        models.AuthType(scheme.Value.Type),
			Name:        name,
			Description: scheme.Value.Description,
		})
	}

	if len(auth) == 0 {
		auth = append(auth, models.AuthScheme{
			Type: models.AuthNone,
		})
	}

	return auth
}
