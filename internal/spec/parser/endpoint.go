package parser

import (
	"strings"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/models"
)

func parseEndpoints(doc *openapi3.T) []models.Endpoint {
	var endpoints []models.Endpoint

	if doc.Paths == nil {
		return endpoints
	}

	for path, pathItem := range doc.Paths.Map() {
		if pathItem == nil {
			continue
		}

		for method, op := range pathItem.Operations() {
			if op == nil {
				continue
			}

			endpoints = append(endpoints, models.Endpoint{
				ID:          op.OperationID,
				Method:      models.HTTPMethod(strings.ToUpper(method)),
				Path:        path,
				Summary:     op.Summary,
				Description: op.Description,
				Deprecated:  op.Deprecated,
				Tags:        op.Tags,
				Parameters:  parseParameters(op),
				RequestBody: parseRequestBody(op),
				Responses:   parseResponses(op),
			})
		}
	}

	return endpoints
}

func parseParameters(op *openapi3.Operation) []models.Parameter {
	var params []models.Parameter

	for _, p := range op.Parameters {
		param := p.Value
		params = append(params, models.Parameter{
			Name:        param.Name,
			In:          models.ParamIn(param.In),
			Required:    param.Required,
			Description: param.Description,
			SchemaRef:   "", // Can be filled in with param.Schema.Ref if needed
		})
	}

	return params
}

func parseRequestBody(op *openapi3.Operation) *models.RequestBody {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}

	rb := op.RequestBody.Value
	return &models.RequestBody{
		Description: rb.Description,
		Required:    rb.Required,
		SchemaRef:   "", // For now, we can extract the schema ref from rb.Content["application/json"].Schema.Ref
	}
}

func parseResponses(op *openapi3.Operation) []models.Response {
	var responses []models.Response

	if op.Responses == nil {
		return responses
	}

	for code, respRef := range op.Responses.Map() {
		if respRef == nil || respRef.Value == nil {
			continue
		}

		description := ""
		if respRef.Value.Description != nil {
			description = *respRef.Value.Description
		}

		responses = append(responses, models.Response{
			StatusCode:  parseStatusCode(code),
			Description: description,
			SchemaRef:   "",
		})
	}

	return responses
}

func parseStatusCode(code string) int {
	// OpenAPI uses strings for status codes
	switch code {
	case "default":
		return 0
	default:
		// convert string to int safely
		if n, err := strconv.Atoi(code); err == nil {
			return n
		}
		return 0
	}
}

