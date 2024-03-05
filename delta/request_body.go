package delta

func calcScoreRequestBody(gt endpoints, spec endpoints) float64 {

	total := len(gt) // ?

	found := 0
	for endpointName, gtEndpoint := range gt {
		if !endpointExists(endpointName, spec) {
			continue
		}
		specEndpoint := spec[endpointName]

		if gtEndpoint.RequestBody == nil && specEndpoint.RequestBody == nil {
			found++
		}
		if gtEndpoint.RequestBody != nil && specEndpoint.RequestBody != nil {
			found++
		}
	}

	added := 0
	for specEndpointName, specEndpoint := range spec {
		if !endpointExists(specEndpointName, gt) {
			continue
		}
		gtEndpoint := gt[specEndpointName]
		if specEndpoint.RequestBody != nil && gtEndpoint.RequestBody == nil {
			added++
		}
	}

	return calcScore(total, found, added)
}

// Total: The total number of request bodies in the label. (every request body has a Required field)
// Removed: A request body required field present in a labeled endpoint but missing from the corresponding endpoint in the result.
// Added: ?
func calcScoreRequestBodyRequired(gt endpoints, spec endpoints) float64 {
	total := 0
	for _, gtEndpoint := range gt {
		if gtEndpoint.RequestBody != nil {
			total++
		}
	}
	found := 0
	for endpointName, gtEndpoint := range gt {
		if !endpointExists(endpointName, spec) {
			continue
		}
		gtRequestBody := gtEndpoint.RequestBody
		if gtRequestBody == nil {
			continue
		}
		specEndpoint := spec[endpointName]
		specRequestBody := specEndpoint.RequestBody
		if specRequestBody == nil {
			continue
		}
		if gtRequestBody.Required == specRequestBody.Required {
			found++
		}
	}
	return calcScore(total, found, 0)
}

// Total: The total number of contents in request bodies.
// Removed: A request body content type field present in a labeled endpoint but missing from the corresponding endpoint in the result.
// Added: ?

func calcScoreRequestBodyContents(gt endpoints, spec endpoints) float64 {
	total := 0
	for _, v := range gt {
		if v.RequestBody == nil {
			continue
		}
		total += len(v.RequestBody.Contents)
	}

	found := 0
	for gtEndpointName, gtEndpoint := range gt {
		if !endpointExists(gtEndpointName, spec) {
			continue
		}
		gtRequestBody := gtEndpoint.RequestBody
		if gtRequestBody == nil {
			continue
		}
		specRequestBody := spec[gtEndpointName].RequestBody
		if specRequestBody == nil {
			continue
		}
		for gtContentType, gtContent := range gtRequestBody.Contents {
			specContent, exists := specRequestBody.Contents[gtContentType]
			if exists && isContentEqual(gtContent, specContent) {
				found++
			}
		}
	}

	// added := 0
	// for specEndpointName, specEndpoint := range spec {
	// 	if !endpointExists(specEndpointName, gt) {
	// 		continue
	// 	}
	// 	specRequestBody := specEndpoint.RequestBody
	// 	if specRequestBody == nil {
	// 		continue
	// 	}
	// 	gtRequestBody := gt[specEndpointName].RequestBody
	// 	if gtRequestBody == nil {
	// 		continue
	// 	}
	// 	for specContentType, specContent := range specRequestBody.Contents {
	// 		gtContent, exists := gtRequestBody.Contents[specContentType]
	// 		if exists && !isContentEqual(gtContent, specContent) {
	// 			added++
	// 		}
	// 	}
	// }

	return calcScore(total, found, 0)
}
