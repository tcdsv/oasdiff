package delta_test

import (
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta"
)

func TestCalcScore_ResponseFound(t *testing.T) {
	weights := delta.NewWeights()
	weights.Responses = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {},
			},
		},
	}

	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ResponseRemoved(t *testing.T) {
	weights := delta.NewWeights()
	weights.Responses = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ResponseAdded(t *testing.T) {
	weights := delta.NewWeights()
	weights.Responses = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusAccepted): {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {},
			},
		},
	}

	require.Equal(t, -1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ResponseAddedEmptyGt(t *testing.T) {
	weights := delta.NewWeights()
	weights.Responses = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {},
			},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ResponseContentFound(t *testing.T) {
	weights := delta.NewWeights()
	weights.ResponsesContent = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {
					Content: map[string]delta.Content{
						delta.ContentTypeJSON: {
							Schema: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {
					Content: map[string]delta.Content{
						delta.ContentTypeJSON: {
							Schema: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
		},
	}

	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ResponseContentNotFound(t *testing.T) {
	weights := delta.NewWeights()
	weights.ResponsesContent = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {
					Content: map[string]delta.Content{
						delta.ContentTypeJSON: {
							Schema: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Responses: map[string]delta.Response{
				delta.GetStatusKey(http.StatusOK): {
					Content: map[string]delta.Content{
						delta.ContentTypeJSON: {
							Schema: &openapi3.Schema{
								Type: "object",
							},
						},
					},
				},
			},
		},
	}

	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}
