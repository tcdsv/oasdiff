package delta

import (
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
)

func Build(spec *openapi3.T) (Endpoints, error) {
	ep := make(Endpoints)

	for path, v := range spec.Paths.Map() {
		basepath, err := spec.Servers.BasePath()

		if err != nil {
			return nil, fmt.Errorf("failed to create endpoints. error: %s", err.Error())
		}

		if basepath != "/" {
			path = basepath + path
		}
		
		extractPath(path, *v, ep)
	}

	return ep, nil
}

func extractPath(path string, pathItem openapi3.PathItem, ep Endpoints) {
	extractOperation(path, pathItem.Get, http.MethodGet, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Head, http.MethodHead, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Post, http.MethodPost, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Put, http.MethodPut, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Patch, http.MethodPatch, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Delete, http.MethodDelete, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Connect, http.MethodConnect, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Options, http.MethodOptions, ep, pathItem.Parameters)
	extractOperation(path, pathItem.Trace, http.MethodTrace, ep, pathItem.Parameters)
}

func extractOperation(path string, op *openapi3.Operation, opName string, ep Endpoints, parameters openapi3.Parameters) {
	if op == nil {
		return
	}
	endpoint := GetPathKey(opName, path)

	ep[endpoint] = Endpoint{
		Parameters: make(map[string]Parameter),
		Responses:  make(map[string]Response),
	}

	params := append(parameters, op.Parameters...)
	extractOperationParameters(endpoint, params, ep)
	extractRequestBody(endpoint, op.RequestBody, ep)
	extractResponses(endpoint, op.Responses, ep)
}

func extractResponses(endpoint string, resp *openapi3.Responses, ep Endpoints) {
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

func extractRequestBody(endpoint string, rb *openapi3.RequestBodyRef, ep Endpoints) {
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

func extractOperationParameters(endpoint string, p openapi3.Parameters, ep Endpoints) {
	if p == nil {
		return
	}
	for _, v := range p {
		if v != nil && v.Value != nil {
			ep[endpoint].Parameters[v.Value.Name] = Parameter{
				Required: v.Value.Required,
			}
		}
	}
}
