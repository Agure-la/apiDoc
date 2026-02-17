package parser

import (
	"github.com/getkin/kin-openapi/openapi3"
    "github.com/agure-la/api-docs/internal/models"

)

// Parser converts OpenAPI specs into domain models.
type Parser struct{}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(apiName string,version string,doc *openapi3.T,) (*spec.APIVersion, error) {

	endpoints := parseEndpoints(doc)
	schemas := parseSchemas(doc)
	auth := parseAuth(doc)

	return &spec.APIVersion{
		Version:   version,
		Info: spec.VersionInfo{
			Title:       doc.Info.Title,
			Description: doc.Info.Description,
		},
		Endpoints: endpoints,
		Schemas:   schemas,
		Auth:      auth,
	}, nil
}
