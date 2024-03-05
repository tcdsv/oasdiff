package delta

func calcScoreResponses(gt endpoints, spec endpoints) float64 {
	total := 0
	for _, gtEndpoint := range gt {
		total += len(gtEndpoint.Responses)
	}

	found := 0
	for gtEndpointName, gtEndpoint := range gt {
		if !endpointExists(gtEndpointName, spec) {
			continue
		}
		specResponses := spec[gtEndpointName].Responses
		for gtResponseCode, _ := range gtEndpoint.Responses {
			_, exists := specResponses[gtResponseCode]
			if exists {
				found++
			}
		}
	}

	added := 0
	for specEndpointName, specEndpoint := range spec {
		if !endpointExists(specEndpointName, spec) {
			continue
		}
		gtResponses := gt[specEndpointName].Responses
		for specResponseCode, _ := range specEndpoint.Responses {
			_, exists := gtResponses[specResponseCode]
			if !exists {
				added++
			}
		}
	}

	return calcScore(total, found, added)
}

func calcScoreResponsesContent(gt endpoints, spec endpoints) float64 {
	total := 0
	for _, endpoint := range gt {
		for _, response := range endpoint.Responses {
			total += len(response.Content)
		}
	}

	found := 0
	for gtEndpointName, gtEndpoint := range gt {
		if !endpointExists(gtEndpointName, spec) {
			continue
		}

		specResponses := spec[gtEndpointName].Responses
		for gtResponseCode, gtResponse := range gtEndpoint.Responses {
			specResponse, exists := specResponses[gtResponseCode]
			if !exists {
				continue
			}

			for contentType, gtContent := range gtResponse.Content {
				specContent, exists := specResponse.Content[contentType]
				if exists && isContentEqual(gtContent, specContent) {
					found++
				}
			}
		}
	}

	return calcScore(total, found, 0)
}
