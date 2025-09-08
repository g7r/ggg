package schema

import (
	"encoding/json"
	"sort"
)

type FieldConfigColorMode string

const (
	FieldConfigColorModeSingleColor    FieldConfigColorMode = "fixed"
	FieldConfigColorModeFromThresholds FieldConfigColorMode = "thresholds"
	FieldConfigColorModeClassicPalette FieldConfigColorMode = "palette-classic"
	FieldConfigColorModeGreenYellowRed FieldConfigColorMode = "continuous-GrYlRd"
	FieldConfigColorModeRedYellowGreen FieldConfigColorMode = "continuous-RdYlGr"
	FieldConfigColorModeBlueYellowRed  FieldConfigColorMode = "continuous-BlYlRd"
	FieldConfigColorModeYellowRed      FieldConfigColorMode = "continuous-YlRd"
	FieldConfigColorModeBluePurple     FieldConfigColorMode = "continuous-BlPu"
	FieldConfigColorModeYellowBlue     FieldConfigColorMode = "continuous-YlBl"
	FieldConfigColorModeBlues          FieldConfigColorMode = "continuous-blues"
	FieldConfigColorModeReds           FieldConfigColorMode = "continuous-reds"
	FieldConfigColorModeGreens         FieldConfigColorMode = "continuous-greens"
	FieldConfigColorModePurples        FieldConfigColorMode = "continuous-purples"
)

type FieldConfigThresholdMode string

const (
	FieldConfigThresholdModeAbsolute   FieldConfigThresholdMode = "absolute"
	FieldConfigThresholdModePercentage FieldConfigThresholdMode = "percentage"
)

type FieldMatcherType string

const (
	FieldMatcherTypeByName       FieldMatcherType = "byName"
	FieldMatcherTypeByRegexp     FieldMatcherType = "byRegexp"
	FieldMatcherTypeByType       FieldMatcherType = "byType"
	FieldMatcherTypeByFrameRefID FieldMatcherType = "byFrameRefID"
)

type DashboardPanelFieldConfig[C any] struct {
	CustomJSON customJSON `json:"-"`

	Defaults  DashboardPanelFieldConfigDefaults[C] `json:"defaults"`
	Overrides []DashboardPanelFieldConfigOverride  `json:"overrides"`
}

type DashboardPanelFieldConfigDefaults[C any] struct {
	CustomJSON customJSON `json:"-"`

	Unit       *Unit                                `json:"unit,omitempty"`
	Color      *DashboardPanelFieldConfigColor      `json:"color,omitempty"`
	Decimals   *int                                 `json:"decimals,omitempty"`
	Min        *float64                             `json:"min,omitempty"`
	Max        *float64                             `json:"max,omitempty"`
	Thresholds *DashboardPanelFieldConfigThresholds `json:"thresholds,omitempty"`
	Mappings   []DashboardPanelFieldConfigMapping   `json:"mappings"`

	Custom *C `json:"custom,omitempty"`
}

type DashboardPanelFieldConfigOverride struct {
	Matcher    DashboardPanelFieldConfigOverrideMatcher    `json:"matcher"`
	Properties []DashboardPanelFieldConfigPropertyOverride `json:"properties"`
}

type DashboardPanelFieldConfigPropertyOverride interface {
	isDashboardPanelFieldConfigPropertyOverride()
}

type DashboardPanelFieldConfigMapping interface {
	isDashboardPanelFieldConfigMapping()
}

type DashboardPanelFieldConfigThresholds struct {
	CustomJSON customJSON `json:"-"`

	Mode  FieldConfigThresholdMode                 `json:"mode"`
	Steps DashboardPanelFieldConfigThresholdsSteps `json:"steps"`
}

type DashboardPanelFieldConfigThresholdsSteps []DashboardPanelFieldConfigThresholdsStep

type DashboardPanelFieldConfigThresholdsStep struct {
	Color string   `json:"color"`
	Value *float64 `json:"value"`
}

type DashboardPanelFieldConfigOverrideMatcher struct {
	ID      FieldMatcherType `json:"id"`
	Options string           `json:"options"`
}

func (d *DashboardPanelFieldConfigThresholds) MarshalJSON() ([]byte, error) {
	type plainThresholds DashboardPanelFieldConfigThresholds
	return marshalResource((*plainThresholds)(d), d.CustomJSON)
}

func (d *DashboardPanelFieldConfigThresholdsSteps) MarshalJSON() ([]byte, error) {
	sortedSteps := make([]DashboardPanelFieldConfigThresholdsStep, len(*d))
	for i := range sortedSteps {
		sortedSteps[i] = (*d)[i]
	}

	sort.Slice(sortedSteps, func(i, j int) bool {
		nilI, nilJ := sortedSteps[i].Value == nil, sortedSteps[j].Value == nil

		if nilI != nilJ {
			// put nil value first
			return nilI
		}

		if !nilI && *sortedSteps[i].Value != *sortedSteps[j].Value {
			return *sortedSteps[i].Value < *sortedSteps[j].Value
		}

		return sortedSteps[i].Color < sortedSteps[j].Color
	})

	return json.Marshal(sortedSteps)
}

type dashboardPanelFieldConfigPropertyOverride struct {
	DashboardPanelFieldConfigPropertyOverride `json:"-"`

	ID    string      `json:"id"`
	Value interface{} `json:"value"`
}

func NewDashboardPanelFieldConfigThresholds() DashboardPanelFieldConfigThresholds {
	return DashboardPanelFieldConfigThresholds{Mode: FieldConfigThresholdModeAbsolute}
}

func NewDashboardPanelFieldConfigPropertyOverrideUnit(unit Unit) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "unit", Value: unit}
}

func NewDashboardPanelFieldConfigPropertyOverrideColor(color *DashboardPanelFieldConfigColor) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "color", Value: color}
}

func NewDashboardPanelFieldConfigPropertyOverrideDecimals(decimals int) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "decimals", Value: decimals}
}

func NewDashboardPanelFieldConfigPropertyOverrideMin(min float64) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "min", Value: min}
}

func NewDashboardPanelFieldConfigPropertyOverrideMax(max float64) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "max", Value: max}
}

func NewDashboardPanelFieldConfigPropertyOverrideThresholds(thresholds *DashboardPanelFieldConfigThresholds) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "thresholds", Value: thresholds}
}

func NewDashboardPanelFieldConfigPropertyOverrideDisplayName(displayName string) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "displayName", Value: displayName}
}

func NewDashboardPanelFieldConfigPropertyOverrideDisplayMode(displayMode string) DashboardPanelFieldConfigPropertyOverride {
	return &dashboardPanelFieldConfigPropertyOverride{ID: "custom.displayMode", Value: displayMode}
}

func (d *DashboardPanelFieldConfig[C]) MarshalJSON() ([]byte, error) {
	type plainConfig DashboardPanelFieldConfig[C]
	return marshalResource((*plainConfig)(d), d.CustomJSON)
}

func (d *DashboardPanelFieldConfigDefaults[C]) MarshalJSON() ([]byte, error) {
	type plainDefaults DashboardPanelFieldConfigDefaults[C]
	return marshalResource((*plainDefaults)(d), d.CustomJSON)
}
