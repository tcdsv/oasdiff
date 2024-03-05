package delta

func calcScoreParams(gt endpoints, spec endpoints) float64 {

	total := 0
	for _, value := range gt {
		total += len(value.Parameters)
	}

	removed := 0
	for endpoint, value := range gt {
		for param, _ := range value.Parameters {
			if _, exists := spec[endpoint].Parameters[param]; !exists {
				removed++
			}
		}
	}

	added := 0
	for endpoint, value := range spec {
		for param, _ := range value.Parameters {
			if _, exists := gt[endpoint].Parameters[param]; !exists {
				added++
			}
		}
	}

	return calcScore(total, total-removed, added)
}

// Total: The total number of parameters in the label.
// Removed: A parameter required field present in a labeled endpoint but missing from the corresponding endpoint in the result.
// Added: ?
func calcScoreParamsRequired(gt endpoints, spec endpoints) float64 {
	total := 0
	for _, value := range gt {
		total += len(value.Parameters)
	}

	removed := 0
	for gtEndpointName, gtEndpoint := range gt {
		if !endpointExists(gtEndpointName, spec) {
			continue
		}

		for gtParamName, gtParamerter := range gtEndpoint.Parameters {
			if !spec[gtEndpointName].hasParameter(gtParamName) {
				// if !parameterExists(gtParamName, spec[gtEndpointName].Parameters) {
				// continue
				// }
				continue
			}
			specParameters := spec[gtEndpointName].Parameters
			specParameter := specParameters[gtParamName]
			if gtParamerter.Required != specParameter.Required {
				removed++
			}
		}
	}

	if total == 0 {
		return 0.0
	}
	return calcScore(total, total-removed, 0)
}
