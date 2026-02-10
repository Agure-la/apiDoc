package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/spec"

)

func parseAuth(doc *openapi3.T) []spec.AuthScheme {
	var auth []spec.AuthScheme

	for name, scheme := range doc.Components.SecuritySchemes {
		auth = append(auth, spec.AuthScheme{
			Type:        spec.AuthType(scheme.Value.Type),
			Name:        name,
			Description: scheme.Value.Description,
		})
	}

	if len(auth) == 0 {
		auth = append(auth, spec.AuthScheme{
			Type: spec.AuthNone,
		})
	}

	return auth
}
