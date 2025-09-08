package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTimeseriesFieldConfigBuilder struct {
	fieldConfig *schema.DashboardPanelFieldConfig[schema.DashboardPanelTimeseriesCustom]
}

type DashboardPanelTimeseriesFieldConfigDefaultsBuilder struct {
	dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelTimeseriesCustom]
	custom *schema.DashboardPanelTimeseriesCustom
}

func (b *DashboardPanelTimeseriesFieldConfigBuilder) Defaults(fn func(builder *DashboardPanelTimeseriesFieldConfigDefaultsBuilder)) {
	fn(&DashboardPanelTimeseriesFieldConfigDefaultsBuilder{
		dashboardPanelFieldConfigDefaultsBuilder: dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelTimeseriesCustom]{
			defaults: &b.fieldConfig.Defaults,
		},
		custom: b.fieldConfig.Defaults.Custom,
	})
}

func (b *DashboardPanelTimeseriesFieldConfigBuilder) AddOverride(fn func(builder *DashboardPanelFieldConfigOverrideBuilder)) {
	var override schema.DashboardPanelFieldConfigOverride
	fn(&DashboardPanelFieldConfigOverrideBuilder{override: &override})
	b.fieldConfig.Overrides = append(b.fieldConfig.Overrides, override)
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) DrawStyle(drawStyle schema.GraphDrawStyle) {
	b.custom.DrawStyle = drawStyle
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) LineInterpolation(lineInterpolation schema.LineInterpolation) {
	b.custom.LineInterpolation = lineInterpolation
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) BarAlignment(alignment schema.BarAlignment) {
	b.custom.BarAlignment = alignment
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) BarWidthFactor(widthFactor float64) {
	b.custom.BarWidthFactor = &widthFactor
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) LineWidth(width float64) {
	b.custom.LineWidth = width
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) FillOpacity(opacity float64) {
	b.custom.FillOpacity = opacity
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) GradientMode(gradientMode schema.GraphGradientMode) {
	b.custom.GradientMode = gradientMode
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) SpanNulls(spanNulls bool) {
	b.custom.SpanNulls = spanNulls
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) InsertNulls(insertNulls bool) {
	b.custom.InsertNulls = insertNulls
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) ShowPoints(visibilityMode schema.VisibilityMode) {
	b.custom.ShowPoints = visibilityMode
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) Stacking(mode schema.StackingMode, group string) {
	b.custom.Stacking.Mode = mode
	b.custom.Stacking.Group = group
}

func (b *DashboardPanelTimeseriesFieldConfigBuilder) JSON(customJSON interface{}) {
	b.fieldConfig.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) DefaultsJSON(customJSON interface{}) {
	b.defaults.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTimeseriesFieldConfigDefaultsBuilder) CustomJSON(customJSON interface{}) {
	b.custom.CustomJSON.Add(customJSON)
}
