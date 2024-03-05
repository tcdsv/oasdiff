package delta

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

func Build(spec *openapi3.T) endpoints {
	ep := make(endpoints)

	for path, v := range spec.Paths.Map() {
		extractPath(path, *v, ep)
	}

	return ep
}

func extractPath(path string, pathItem openapi3.PathItem, ep endpoints) {
	extractOperation(path, pathItem.Get, http.MethodGet, ep)
	extractOperation(path, pathItem.Head, http.MethodHead, ep)
	extractOperation(path, pathItem.Post, http.MethodPost, ep)
	extractOperation(path, pathItem.Put, http.MethodPut, ep)
	extractOperation(path, pathItem.Patch, http.MethodPatch, ep)
	extractOperation(path, pathItem.Delete, http.MethodDelete, ep)
	extractOperation(path, pathItem.Connect, http.MethodConnect, ep)
	extractOperation(path, pathItem.Options, http.MethodOptions, ep)
	extractOperation(path, pathItem.Trace, http.MethodTrace, ep)
}

func extractOperation(path string, op *openapi3.Operation, opName string, ep endpoints) {
	if op == nil {
		return
	}
	endpoint := GetPathKey(opName, path)

	ep[endpoint] = Endpoint{
		Parameters: make(map[string]Parameter),
		Responses:  make(map[string]Response),
	}

	extractOperationParameters(endpoint, &op.Parameters, ep)
	extractRequestBody(endpoint, op.RequestBody, ep)
	extractResponses(endpoint, op.Responses, ep)
}

func extractResponses(endpoint string, resp *openapi3.Responses, ep endpoints) {
	if resp == nil {
		return
	}
	responses := resp.Map()
	if len(responses) == 0 {
		return
	}
	for code, r := range responses {
		if r.Value == nil {
			ep[endpoint].Responses[code] = Response{}
			continue
		}

		response := &Response{
			Content: make(map[string]Content),
		}

		contents := r.Value.Content
		for contentType, c := range contents {
			if c.Schema == nil {
				continue
			}
			if c.Schema.Value == nil {
				continue
			}
			response.Content[contentType] = Content{
				Schema: c.Schema.Value,
			}
		}
		ep[endpoint].Responses[code] = *response
	}
}

func extractRequestBody(endpoint string, rb *openapi3.RequestBodyRef, ep endpoints) {
	if rb == nil {
		return
	}
	if rb.Value == nil {
		return
	}
	epoint := ep[endpoint]

	requestBody := &RequestBody{
		Contents: make(map[string]Content),
	}

	for contentType, content := range rb.Value.Content {
		if content == nil {
			continue
		}
		if content.Schema == nil {
			continue
		}
		requestBody.Contents[contentType] = Content{
			Schema: content.Schema.Value,
		}
	}
	requestBody.Required = rb.Value.Required
	epoint.RequestBody = requestBody
	ep[endpoint] = epoint
}

func extractOperationParameters(endpoint string, p *openapi3.Parameters, ep endpoints) {
	if p == nil {
		return
	}
	for _, v := range *p {
		if v != nil && v.Value != nil {
			ep[endpoint].Parameters[v.Value.Name] = Parameter{
				Required: v.Value.Required,
			}
		}
	}
}
