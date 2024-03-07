package delta

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/sirupsen/logrus"
	"github.com/tufin/oasdiff/diff"
)

func isContentEqual(gt Content, spec Content) bool {

	if gt.Schema == nil && spec.Schema == nil {
		return true
	}
	if (gt.Schema == nil && spec.Schema != nil) || (gt.Schema != nil && spec.Schema == nil) {
		return false
	}

	res, err := diff.GetSchemaDiff2(&openapi3.SchemaRef{
		Value: gt.Schema,
	}, &openapi3.SchemaRef{
		Value: spec.Schema,
	})
	if err != nil {
		logrus.Errorf("failed to get schema diff with '%s'", err)
		return false
	}

	return res.Empty()
}
