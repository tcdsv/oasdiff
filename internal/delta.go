package internal

import (
	"fmt"
	"io"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
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
	base := flags.getBase()
	revision := flags.getRevision()

	gt := delta.Build(base)
	spec := delta.Build(revision)
	weights := delta.NewWeights()

	_, _ = fmt.Fprintf(stdout, "%g\n", delta.CalcScore(weights, gt, spec))
	return false, nil
}
