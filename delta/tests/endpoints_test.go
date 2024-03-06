package delta_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta"
)

func TestCalcScore_EndpointsFound(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.Endpoints = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}
	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}
	require.Equal(t, 1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_EndpointsAdded(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.Endpoints = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}
	spec := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodPut, "/path"): {},
	}
	require.Equal(t, -1.0, delta.CalcScore(weights, gt, spec))
}

func TestCalcScore_EndpointsRemoved(t *testing.T) {

	weights := delta.EmptyWeights()
	weights.Endpoints = 1.0

	gt := map[string]delta.Endpoint{
		delta.GetPathKey(http.MethodGet, "/path"): {},
	}
	spec := map[string]delta.Endpoint{}
	require.Equal(t, 0.0, delta.CalcScore(weights, gt, spec))
}
