package parser

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/agure-la/api-docs/internal/spec"
     "strconv"
)

func parseEndpoints(doc *openapi3.T) []spec.Endpoint {
	var endpoints []spec.Endpoint

	for path, pathItem := range *doc.Paths { // <- dereference here
		for method, op := range pathItem.Operations() {
			endpoints = append(endpoints, spec.Endpoint{
				ID:          op.OperationID,
				Method:      spec.HTTPMethod(strings.ToUpper(method)),
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


func parseParameters(op *openapi3.Operation) []spec.Parameter {
	var params []spec.Parameter

	for _, p := range op.Parameters {
		param := p.Value
		params = append(params, spec.Parameter{
			Name:        param.Name,
			In:          spec.ParamIn(param.In),
			Required:    param.Required,
			Description: param.Description,
			SchemaRef:   "", // Can be filled in with param.Schema.Ref if needed
		})
	}

	return params
}

func parseRequestBody(op *openapi3.Operation) *spec.RequestBody {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}

	rb := op.RequestBody.Value
	return &spec.RequestBody{
		Description: rb.Description,
		Required:    rb.Required,
		SchemaRef:   "", // For now, we can extract the schema ref from rb.Content["application/json"].Schema.Ref
	}
}

func parseResponses(op *openapi3.Operation) []spec.Response {
	var responses []spec.Response

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

		responses = append(responses, spec.Response{
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

