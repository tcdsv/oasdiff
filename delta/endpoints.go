package delta

type Endpoint struct {
	Parameters  map[string]Parameter
	RequestBody *RequestBody
	Responses   map[string]Response
}

func (e Endpoint) hasParameter(parameter string) bool {
	_, exists := e.Parameters[parameter]
	return exists
}

func (e Endpoint) hasRequestBody() bool {
	return e.RequestBody != nil
}

func (e Endpoint) hasResponse(responseCode string) bool {
	_, exists := e.Parameters[responseCode]
	return exists
}

func calcScoreEndpoints(gt endpoints, spec endpoints) float64 {
	removed := 0
	for key := range gt {
		if _, exists := spec[key]; !exists {
			removed++
		}
	}

	added := 0
	for key := range spec {
		if _, exists := gt[key]; !exists {
			added++
		}
	}

	total := len(gt)
	found := total - removed
	return calcScore(len(gt), found, added)
}
