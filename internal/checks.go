package internal

import (
	"fmt"
	"io"
	"sort"

	"github.com/spf13/cobra"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/checker/localizations"
	"github.com/tufin/oasdiff/formatters"
	"golang.org/x/exp/slices"
)

const checksCmd = "checks"

type ChecksFlags struct {
	lang     string
	format   string
	severity []string
	tags     []string
}

func getChecksCmd() *cobra.Command {
	flags := ChecksFlags{}

	cmd := cobra.Command{
		Use:               "checks [flags]",
		Short:             "Display checks",
		Long:              `Display a list of all supported checks.`,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions, // see https://github.com/spf13/cobra/issues/1969
		RunE: func(cmd *cobra.Command, args []string) error {
			// by now flags have been parsed successfully, so we don't need to show usage on any errors
			cmd.Root().SilenceUsage = true

			if err := runChecks(cmd.OutOrStdout(), flags); err != nil {
				setReturnValue(cmd, err.Code)
				return err
			}

			return nil
		},
	}

	enumWithOptions(&cmd, newEnumValue(localizations.GetSupportedLanguages(), localizations.LangDefault, &flags.lang), "lang", "l", "language for localized output")
	enumWithOptions(&cmd, newEnumValue(formatters.SupportedFormatsByContentType(formatters.OutputChecks), string(formatters.FormatText), &flags.format), "format", "f", "output format")
	enumWithOptions(&cmd, newEnumSliceValue([]string{"info", "warn", "error"}, nil, &flags.severity), "severity", "s", "include only checks with any of specified severities")
	enumWithOptions(&cmd, newEnumSliceValue(getAllTags(), nil, &flags.tags), "tags", "t", "include only checks with all specified tags")

	return &cmd
}

func runChecks(stdout io.Writer, flags ChecksFlags) *ReturnError {
	return outputChecks(stdout, flags, checker.GetAllRules())
}

func outputChecks(stdout io.Writer, flags ChecksFlags, rules []checker.BackwardCompatibilityRule) *ReturnError {
	// formatter lookup
	formatter, err := formatters.Lookup(flags.format, formatters.FormatterOpts{
		Language: flags.lang,
	})
	if err != nil {
		return getErrUnsupportedFormat(flags.format, checksCmd)
	}

	// filter rules
	checks := make(formatters.Checks, 0, len(rules))
	for _, rule := range rules {
		// severity
		if len(flags.severity) > 0 {
			if rule.Level == checker.ERR && !slices.Contains(flags.severity, "error") {
				continue
			}
			if rule.Level == checker.WARN && !slices.Contains(flags.severity, "warn") {
				continue
			}
			if rule.Level == checker.INFO && !slices.Contains(flags.severity, "info") {
				continue
			}
		}

		// tags
		if !matchTags(flags.tags, rule) {
			continue
		}

		checks = append(checks, formatters.Check{
			Id:          rule.Id,
			Level:       rule.Level.String(),
			Description: rule.Description,
		})
	}

	// render
	sort.Sort(checks)
	bytes, err := formatter.RenderChecks(checks, formatters.NewRenderOpts())
	if err != nil {
		return getErrFailedPrint("checks "+flags.format, err)
	}

	// print output
	_, _ = fmt.Fprintf(stdout, "%s\n", bytes)

	return nil
}
