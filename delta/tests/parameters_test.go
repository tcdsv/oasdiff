package delta_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta"
)

func TestCalcScore_ParametersFound(t *testing.T) {

	weights := delta.NewWeights()
	weights.Parameters = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {},
			},
		},
	}
	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersAdded(t *testing.T) {

	weights := delta.NewWeights()
	weights.Parameters = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"id": {},
			},
		},
	}
	require.Equal(t, -1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersRemoved(t *testing.T) {

	weights := delta.NewWeights()
	weights.Parameters = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersAddedEmptyGt(t *testing.T) {

	weights := delta.NewWeights()
	weights.Parameters = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"id": {},
			},
		},
	}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersRequiredFound(t *testing.T) {
	weights := delta.NewWeights()
	weights.ParametersRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {Required: true},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {Required: true},
			},
		},
	}
	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersRequiredFoundAdded(t *testing.T) {

	weights := delta.NewWeights()
	weights.ParametersRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {Required: true},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"id": {Required: true},
			},
		},
	}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersRequiredFoundRemoved(t *testing.T) {

	weights := delta.NewWeights()
	weights.ParametersRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {Required: true},
			},
		},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"name": {Required: false},
			},
		},
	}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_ParametersRequiredFoundAddedEmptyGt(t *testing.T) {

	weights := delta.NewWeights()
	weights.ParametersRequired = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}

	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {
			Parameters: map[string]delta.Parameter{
				"id": {Required: true},
			},
		},
	}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}
