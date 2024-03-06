package delta

import (
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
)

const (
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
	ContentTypePlainText = "text/plain"
)

type Weights struct {
	Endpoints           float64
	Parameters          float64
	ParametersRequired  float64
	RequestBody         float64
	RequestBodyRequired float64
	RequestBodyContent  float64
	Responses           float64
	ResponsesContent    float64
}

type Parameter struct {
	Required bool
}

type RequestBody struct {
	Contents map[string]Content
	Required bool
}

type Content struct {
	Schema *openapi3.Schema
}

type Response struct {
	Content map[string]Content
}

type endpoints map[string]Endpoint

func calcScore(total int, found int, added int) float64 {
	if total == 0 {
		return 0
	}
	point := 1 / float64(total)
	return (float64(found) * point) - (float64(added) * point)
}

func GetPathKey(method string, path string) string {
	return method + "$" + path
}

func GetStatusKey(status int) string {
	return strconv.Itoa(status)
}

func endpointExists(endpoint string, spec endpoints) bool {
	_, exists := spec[endpoint]
	return exists
}

func NewWeights() Weights {
	return Weights{
		Endpoints:           0.0,
		Parameters:          0.0,
		ParametersRequired:  0.0,
		RequestBody:         0.0,
		RequestBodyRequired: 0.0,
		RequestBodyContent:  0.0,
		Responses:           0.0,
		ResponsesContent:    0.0,
	}
}

func CalcScore(weights Weights, gt endpoints, spec endpoints) float64 {
	endpoints := weights.Endpoints * calcScoreEndpoints(gt, spec)
	parameters := weights.Parameters * calcScoreParams(gt, spec)
	parametersRequired := weights.ParametersRequired * calcScoreParamsRequired(gt, spec)
	requestBody := weights.RequestBody * calcScoreRequestBody(gt, spec)
	requestBodyRequired := weights.RequestBodyRequired * calcScoreRequestBodyRequired(gt, spec)
	requestBodyContent := weights.RequestBodyContent * calcScoreRequestBodyContents(gt, spec)
	responses := weights.Responses * calcScoreResponses(gt, spec)
	responsesContent := weights.ResponsesContent * calcScoreResponsesContent(gt, spec)
	return endpoints +
		parameters +
		parametersRequired +
		requestBody +
		requestBodyRequired +
		requestBodyContent +
		responses +
		responsesContent
}
