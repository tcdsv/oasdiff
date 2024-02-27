package delta2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta2"
	"github.com/tufin/oasdiff/diff"
)

func Test_NoDifference(t *testing.T) {
	label := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
		{Path: "/abc", Method: "POST"},
	}

	generated := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
		{Path: "/abc", Method: "POST"},
	}

	require.Equal(t, float64(1), delta2.Get(label, generated))
}

func Test_Difference(t *testing.T) {
	label := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
		{Path: "/abc", Method: "POST"},
	}

	generated := []diff.Endpoint{}
	require.Equal(t, float64(0), delta2.Get(label, generated))
}

func Test_Missing(t *testing.T) {
	label := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
		{Path: "/abc", Method: "POST"},
	}

	generated := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
	}

	require.Equal(t, float64(0.5), delta2.Get(label, generated))
}

func Test_Wrong(t *testing.T) {
	label := []diff.Endpoint{
		{Path: "/abc", Method: "GET"},
		{Path: "/abc", Method: "POST"},
	}

	generated := []diff.Endpoint{
		{Path: "/abc2", Method: "GET"},
		{Path: "/abc3", Method: "GET"},
		{Path: "/abc4", Method: "GET"},
		{Path: "/abc5", Method: "GET"},
	}

	require.Equal(t, float64(-2), delta2.Get(label, generated))
}

func Test_Wrong4(t *testing.T) {
	label := []diff.Endpoint{
		{Path: "/users", Method: "GET"},
		{Path: "/users/{userId}", Method: "GET"},
		{Path: "/products", Method: "GET"},
	}

	generated := []diff.Endpoint{
		{Path: "/users", Method: "GET"},
		{Path: "/products", Method: "GET"},
	}

	require.Equal(t, 2/float64(3), delta2.Get(label, generated))
}
