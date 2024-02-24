package internal

import (
	"fmt"
	"io"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"github.com/tufin/oasdiff/delta2"
	"github.com/tufin/oasdiff/diff"
)

func getDeltaCmd() *cobra.Command {

	flags := DeltaFlags{}

	cmd := cobra.Command{
		Use:   "delta base revision [flags]",
		Short: "Calculate the delta value",
		Long:  `Calculate a numeric value representing the delta between base and revision specs.` + specHelp,
		Args:  getParseArgs(&flags),
		RunE:  getRun(&flags, runDelta),
	}

	addCommonDiffFlags(&cmd, &flags)
	// cmd.PersistentFlags().BoolVarP(&flags.asymmetric, "asymmetric", "", false, "perform asymmetric diff (only elements of base that are missing in revision)")

	return &cmd
}

func runDelta(flags Flags, stdout io.Writer) (bool, *ReturnError) {

	openapi3.CircularReferenceCounter = flags.getCircularReferenceCounter()

	diffResult, err := calcDiff(flags)
	if err != nil {
		return false, err
	}

	labelEndpoints := getEndpoints(diffResult.specInfoPair.Base.Spec)
	generatedEndpoints := getEndpoints(diffResult.specInfoPair.Revision.Spec)
	result := delta2.Get(labelEndpoints, generatedEndpoints)
	_, _ = fmt.Fprintf(stdout, "score: %g\ndiscovered: %v\nmissing: %v\nwrong: %v\n", result.Score, result.Discovered, result.Missing, result.Wrong)

	// _, _ = fmt.Fprintf(stdout, "%g\n", delta.Get(flags.getAsymmetric(), diffResult.diffReport))

	return false, nil
}

func getEndpoints(labelSpec *openapi3.T) []diff.Endpoint {
	endpoints := []diff.Endpoint{}

	var operations = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	for path, v := range labelSpec.Paths.Map() {

		for _, op := range operations {
			opItem := v.GetOperation(op)
			if opItem != nil {
				endpoints = append(endpoints, diff.Endpoint{
					Path:   path,
					Method: op,
				})
			}
		}

	}

	return endpoints
}
