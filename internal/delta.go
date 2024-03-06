package internal

import (
	"fmt"
	"io"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"github.com/tufin/oasdiff/delta"
	"github.com/tufin/oasdiff/load"
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

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true

	flattenAllOf := load.GetOption(load.WithFlattenAllOf(), flags.getFlattenAllOf())
	flattenParams := load.GetOption(load.WithFlattenParams(), flags.getFlattenParams())
	lowerHeaderNames := load.GetOption(load.WithLowercaseHeaders(), flags.getInsensitiveHeaders())

	s1, err := load.NewSpecInfo(loader, flags.getBase(), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return false, getErrFailedToLoadSpec("base", flags.getBase(), err)
	}

	s2, err := load.NewSpecInfo(loader, flags.getRevision(), flattenAllOf, flattenParams, lowerHeaderNames)
	if err != nil {
		return false, getErrFailedToLoadSpec("revision", flags.getRevision(), err)
	}

	gt := delta.Build(s1.Spec)
	spec := delta.Build(s2.Spec)

	weights := delta.DefaultWeights()
	_, _ = fmt.Fprintf(stdout, "%g\n", delta.CalcScore(weights, gt, spec))
	return false, nil
}
