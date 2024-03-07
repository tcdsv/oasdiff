package delta

import (
	"log/slog"
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

type Endpoints map[string]Endpoint

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

func endpointExists(endpoint string, spec Endpoints) bool {
	_, exists := spec[endpoint]
	return exists
}

func DefaultWeights() Weights {
	return Weights{
		Endpoints:           0.5,
		Parameters:          0.5 / 7,
		ParametersRequired:  0.5 / 7,
		RequestBody:         0.5 / 7,
		RequestBodyRequired: 0.5 / 7,
		RequestBodyContent:  0.5 / 7,
		Responses:           0.5 / 7,
		ResponsesContent:    0.5 / 7,
	}
}

func EmptyWeights() Weights {
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

func CalcScoreFiles(basePath string, revisionPath string) (float64, error) {

	sl := openapi3.NewLoader()

	baseSpec, err := sl.LoadFromFile(basePath)
	if err != nil {
		slog.Error("failed to load base spec", "error", err)
		return -1, err
	}

	revisionSpec, err := sl.LoadFromFile(revisionPath)
	if err != nil {
		slog.Error("failed to load revision spec", "error", err)
		return -1, err
	}

	return CalcScore(DefaultWeights(), Build(baseSpec), Build(revisionSpec)), nil
}

func CalcScore(weights Weights, gt Endpoints, spec Endpoints) float64 {
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
