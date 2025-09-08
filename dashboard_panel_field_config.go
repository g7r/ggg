package ggg

import "github.com/g7r/ggg/schema"

type dashboardPanelFieldConfigDefaultsBuilder[C any] struct {
	defaults *schema.DashboardPanelFieldConfigDefaults[C]
}

type DashboardPanelFieldConfigOverrideBuilder struct {
	override *schema.DashboardPanelFieldConfigOverride
}

type DashboardPanelFieldConfigThresholdsBuilder struct {
	thresholds *schema.DashboardPanelFieldConfigThresholds
}

type DashboardPanelFieldConfigColorBuilder struct {
	color *schema.DashboardPanelFieldConfigColor
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Color(fn func(builder *DashboardPanelFieldConfigColorBuilder)) {
	var color schema.DashboardPanelFieldConfigColor
	fn(&DashboardPanelFieldConfigColorBuilder{color: &color})
	b.defaults.Color = &color
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Unit(unit schema.Unit) {
	b.defaults.Unit = &unit
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Min(min float64) {
	b.defaults.Min = &min
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Max(max float64) {
	b.defaults.Max = &max
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Decimals(decimals int) {
	b.defaults.Decimals = &decimals
}

func (b *dashboardPanelFieldConfigDefaultsBuilder[C]) Thresholds(fn func(builder *DashboardPanelFieldConfigThresholdsBuilder)) {
	thresholds := schema.NewDashboardPanelFieldConfigThresholds()
	fn(&DashboardPanelFieldConfigThresholdsBuilder{thresholds: &thresholds})
	b.defaults.Thresholds = &thresholds
}

func (b *DashboardPanelFieldConfigOverrideBuilder) Matcher(matcherType schema.FieldMatcherType, options string) {
	b.override.Matcher = schema.DashboardPanelFieldConfigOverrideMatcher{ID: matcherType, Options: options}
}

func (b *DashboardPanelFieldConfigOverrideBuilder) AddDisplayNameOverride(displayName string) {
	b.override.Properties = append(b.override.Properties, schema.NewDashboardPanelFieldConfigPropertyOverrideDisplayName(displayName))
}

func (b *DashboardPanelFieldConfigOverrideBuilder) AddUnitOverride(unit schema.Unit) {
	b.override.Properties = append(b.override.Properties, schema.NewDashboardPanelFieldConfigPropertyOverrideUnit(unit))
}

func (b *DashboardPanelFieldConfigOverrideBuilder) AddDecimalsOverride(decimals int) {
	b.override.Properties = append(b.override.Properties, schema.NewDashboardPanelFieldConfigPropertyOverrideDecimals(decimals))
}

func (b *DashboardPanelFieldConfigOverrideBuilder) AddThresholdsOverride(fn func(builder *DashboardPanelFieldConfigThresholdsBuilder)) {
	thresholds := schema.NewDashboardPanelFieldConfigThresholds()
	fn(&DashboardPanelFieldConfigThresholdsBuilder{thresholds: &thresholds})
	b.override.Properties = append(b.override.Properties, schema.NewDashboardPanelFieldConfigPropertyOverrideThresholds(&thresholds))
}

func (b *DashboardPanelFieldConfigOverrideBuilder) AddDisplayModeOverride(displayMode string) {
	b.override.Properties = append(b.override.Properties, schema.NewDashboardPanelFieldConfigPropertyOverrideDisplayMode(displayMode))
}

func (b *DashboardPanelFieldConfigThresholdsBuilder) Mode(mode schema.FieldConfigThresholdMode) {
	b.thresholds.Mode = mode
}

func (b *DashboardPanelFieldConfigThresholdsBuilder) AddDefaultStep(color string) {
	b.thresholds.Steps = append(b.thresholds.Steps, schema.DashboardPanelFieldConfigThresholdsStep{Color: color})
}

func (b *DashboardPanelFieldConfigThresholdsBuilder) AddStep(threshold float64, color string) {
	b.thresholds.Steps = append(b.thresholds.Steps, schema.DashboardPanelFieldConfigThresholdsStep{Value: &threshold, Color: color})
}

func (b *DashboardPanelFieldConfigColorBuilder) FixedColor(color string) {
	b.color.FixedColor = color
}

func (b *DashboardPanelFieldConfigColorBuilder) Mode(mode schema.FieldConfigColorMode) {
	b.color.Mode = mode
}

func (b *DashboardPanelFieldConfigThresholdsBuilder) JSON(customJSON interface{}) {
	b.thresholds.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelFieldConfigColorBuilder) JSON(customJSON interface{}) {
	b.color.CustomJSON.Add(customJSON)
}
