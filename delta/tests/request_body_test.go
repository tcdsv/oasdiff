package delta_test

import (
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta"
)

func TestCalcScore_RequestBodyFound(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.RequestBody = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{},
		},
	}

	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyAdded(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.RequestBody = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: nil,
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{},
		},
	}

	require.Equal(t, -1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyRemoved(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.RequestBody = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: nil,
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyRequiredFound(t *testing.T) {
	weights := delta.EmptyWeights()
	weights.RequestBodyRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Required: false,
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Required: false,
			},
		},
	}

	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyRequiredRemoved(t *testing.T) {
	weights := delta.EmptyWeights()
	weights.RequestBodyRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Required: false,
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Required: true,
			},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyRequiredAdded(t *testing.T) {
	weights := delta.EmptyWeights()
	weights.RequestBodyRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: nil,
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Required: true,
			},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyContentSchemaFound(t *testing.T) {
	weights := delta.EmptyWeights()
	weights.RequestBodyContent = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Contents: map[string]delta.Content{
					delta.ContentTypeJSON: {
						Schema: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Contents: map[string]delta.Content{
					delta.ContentTypeJSON: {
						Schema: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	}

	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_RequestBodyContentSchemaNotFound(t *testing.T) {
	weights := delta.EmptyWeights()
	weights.RequestBodyContent = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Contents: map[string]delta.Content{
					delta.ContentTypeJSON: {
						Schema: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			RequestBody: &delta.RequestBody{
				Contents: map[string]delta.Content{
					delta.ContentTypeJSON: {
						Schema: &openapi3.Schema{
							Type: "object",
						},
					},
				},
			},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}
