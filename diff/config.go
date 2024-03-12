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
}

const (
	ExcludeExamplesOption             = "examples"
	ExcludeDescriptionOption          = "description"
	ExcludeEndpointsOption            = "endpoints"
	ExcludeTitleOption                = "title"
	ExcludeSummaryOption              = "summary"
	ExcludeFormatOption               = "format"
	ExcludeDefaultOption              = "default"
	AdditionalPropertiesAllowedOption = "additionalPropsAllowed"
	UniqueItemsOption                 = "uniqueItems"
	ExclusiveMinOption                = "exclusiveMin"
	ExclusiveMaxOption                = "exclusiveMax"
	NullableOption                    = "nullable"
	ReadOnlyOption                    = "readonly"
	WriteOnlyOption                   = "writeonly"
	AllowEmptyValueOption             = "allowEmptyValue"
	XMLOption                         = "xml"
	DeprecatedOption                  = "deprecated"
	MinOption                         = "min"
	MaxOption                         = "max"
	MultipleOfOption                  = "multipleOf"
	MinLengthOption                   = "minLength"
	MaxLengthOption                   = "maxLength"
	PatternOption                     = "pattern"
	MinItemsOption                    = "minItems"
	MaxItemsOption                    = "maxItems"
	MinPropsDiff                      = "minProps"
	MaxPropsDiff                      = "maxProps"
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

func (config *Config) IsExcludeFormat() bool {
	return config.ExcludeElements.Contains(ExcludeFormatOption)
}

func (config *Config) IsExcludeDefault() bool {
	return config.ExcludeElements.Contains(ExcludeDefaultOption)
}

func (config *Config) IsExcludeAdditionalPropertiesAllowed() bool {
	return config.ExcludeElements.Contains(AdditionalPropertiesAllowedOption)
}

func (config *Config) IsExcludeUniqueItems() bool {
	return config.ExcludeElements.Contains(UniqueItemsOption)
}

func (config *Config) IsExcludeExclusiveMin() bool {
	return config.ExcludeElements.Contains(ExclusiveMinOption)
}

func (config *Config) IsExcludeExclusiveMax() bool {
	return config.ExcludeElements.Contains(ExclusiveMaxOption)
}

func (config *Config) IsExcludeNullable() bool {
	return config.ExcludeElements.Contains(NullableOption)
}

func (config *Config) IsExcludeReadOnly() bool {
	return config.ExcludeElements.Contains(ReadOnlyOption)
}

func (config *Config) IsExcludeWriteOnly() bool {
	return config.ExcludeElements.Contains(WriteOnlyOption)
}

func (config *Config) IsExcludeAllowEmptyValue() bool {
	return config.ExcludeElements.Contains(AllowEmptyValueOption)
}

func (config *Config) IsExcludeXML() bool {
	return config.ExcludeElements.Contains(XMLOption)
}

func (config *Config) IsExcludeDeprecated() bool {
	return config.ExcludeElements.Contains(DeprecatedOption)
}

func (config *Config) IsExcludeMin() bool {
	return config.ExcludeElements.Contains(MinOption)
}

func (config *Config) IsExcludeMax() bool {
	return config.ExcludeElements.Contains(MaxOption)
}

func (config *Config) IsExcludeMultipleOf() bool {
	return config.ExcludeElements.Contains(MultipleOfOption)
}

func (config *Config) IsExcludeMinLength() bool {
	return config.ExcludeElements.Contains(MinLengthOption)
}

func (config *Config) IsExcludeMaxLength() bool {
	return config.ExcludeElements.Contains(MaxLengthOption)
}

func (config *Config) IsExcludePattern() bool {
	return config.ExcludeElements.Contains(PatternOption)
}

func (config *Config) IsExcludeMinItems() bool {
	return config.ExcludeElements.Contains(MinItemsOption)
}

func (config *Config) IsExcludeMaxItems() bool {
	return config.ExcludeElements.Contains(MaxItemsOption)
}

func (config *Config) IsExcludeMinProps() bool {
	return config.ExcludeElements.Contains(MinPropsDiff)
}

func (config *Config) IsExcludeMaxProps() bool {
	return config.ExcludeElements.Contains(MaxPropsDiff)
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
