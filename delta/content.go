package delta

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/tufin/oasdiff/diff"
)

func isContentEqual(gt Content, spec Content) bool {
	gtRef := &openapi3.SchemaRef{
		Value: gt.Schema,
	}
	resultRef := &openapi3.SchemaRef{
		Value: spec.Schema,
	}
	res, _ := diff.GetSchemaDiff2(gtRef, resultRef)
	return res.Empty()
}
