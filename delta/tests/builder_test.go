package delta_test

import (
	"context"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/delta"
)

func Test_ItBuildsSpec(t *testing.T) {
	spec := loadSpec(t, "testdata/sample.yaml")
	ep := delta.Build(spec) //todo handle errors.
	require.NotEmpty(t, ep)
}

func loadSpec(t *testing.T, path string) *openapi3.T {
	ctx := context.Background()
	sl := openapi3.NewLoader()
	doc, err := sl.LoadFromFile(path)
	require.NoError(t, err, "loading test file")
	err = doc.Validate(ctx)
	require.NoError(t, err)
	return doc
}
