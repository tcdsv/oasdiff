package delta2

import (
	"fmt"

	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/utils"
)

type Result struct {
	Score          float64
	LabelEndpoints int
	Discovered     utils.StringList
	Removed        utils.StringList
	Added          utils.StringList
}

func Get(labelEndpoints []diff.Endpoint, generatedEndpoints []diff.Endpoint) float64 {
	labelEndpointsStr := endpointToStr(labelEndpoints)
	generatedEndpointsStr := endpointToStr(generatedEndpoints)

	wrong := generatedEndpointsStr.Minus(labelEndpointsStr)
	discovered := generatedEndpointsStr.Intersection(labelEndpointsStr)

	unitscore := float64(1) / float64(len(labelEndpoints))

	return (float64(len(discovered)) * unitscore) - (float64(len(wrong)) * unitscore)
}

func endpointToStr(endpoints []diff.Endpoint) utils.StringSet {
	ep := make(utils.StringSet)
	for _, e := range endpoints {
		ep[fmt.Sprintf(e.Method+" "+e.Path)] = struct{}{}
	}
	return ep
}
