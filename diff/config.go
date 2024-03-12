package diff

import (
	"github.com/tufin/oasdiff/utils"
)

// Config includes various settings to control the diff
type Config struct {
	IncludeExtensions       utils.StringSet
	PathFilter              string
	FilterExtension         string
	PathPrefixBase          string
	PathPrefixRevision      string
	PathStripPrefixBase     string
	PathStripPrefixRevision string
	ExcludeElements         utils.StringSet
	IncludePathParams       bool
	EquivalenElements       utils.StringSet
}

const (
	ExcludeExamplesOption    = "examples"
	ExcludeDescriptionOption = "description"
	ExcludeEndpointsOption   = "endpoints"
	ExcludeTitleOption       = "title"
	ExcludeSummaryOption     = "summary"
	EquivalentFormat         = "format"
)

var ExcludeDiffOptions = []string{
	ExcludeExamplesOption,
	ExcludeDescriptionOption,
	ExcludeEndpointsOption,
	ExcludeTitleOption,
	ExcludeSummaryOption,
}

// NewConfig returns a default configuration
func NewConfig() *Config {
	return &Config{
		IncludeExtensions: utils.StringSet{},
		ExcludeElements:   utils.StringSet{},
		EquivalenElements: utils.StringSet{},
	}
}

func (config *Config) WithExcludeElements(excludeElements []string) *Config {
	config.ExcludeElements = utils.StringList(excludeElements).ToStringSet()
	return config
}

func (config *Config) IsExcludeExamples() bool {
	return config.ExcludeElements.Contains(ExcludeExamplesOption)
}

func (config *Config) IsExcludeDescription() bool {
	return config.ExcludeElements.Contains(ExcludeDescriptionOption)
}

func (config *Config) IsExcludeEndpoints() bool {
	return config.ExcludeElements.Contains(ExcludeEndpointsOption)
}

func (config *Config) IsExcludeTitle() bool {
	return config.ExcludeElements.Contains(ExcludeTitleOption)
}

func (config *Config) IsExcludeSummary() bool {
	return config.ExcludeElements.Contains(ExcludeSummaryOption)
}

func (config *Config) IsEquivalentFormat() bool {
	return config.EquivalenElements.Contains(EquivalentFormat)
} 

const (
	SunsetExtension          = "x-sunset"
	XStabilityLevelExtension = "x-stability-level"
	XExtensibleEnumExtension = "x-extensible-enum"
)

func (config *Config) WithCheckBreaking() *Config {
	config.IncludeExtensions.Add(XStabilityLevelExtension)
	config.IncludeExtensions.Add(SunsetExtension)
	config.IncludeExtensions.Add(XExtensibleEnumExtension)

	return config
}
