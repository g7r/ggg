package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTimeseriesBuilder struct {
	dashboardPanelBuilder
	panel *schema.DashboardPanelTimeseries

	nextRefID rune
}

type DashboardPanelTimeseriesOptionsBuilder struct {
	parent *DashboardPanelTimeseriesBuilder
}

type DashboardPanelTimeseriesOptionsTooltipBuilder struct {
	parent *DashboardPanelTimeseriesBuilder
}

type DashboardPanelTimeseriesOptionsLegendBuilder struct {
	parent *DashboardPanelTimeseriesBuilder
}

type DashboardPanelTargetsBuilder struct {
	nextRefID *rune
	targets   *[]schema.DashboardPanelGenericTarget
}

func newDashboardPanelTimeseriesBuilder() *DashboardPanelTimeseriesBuilder {
	panel := schema.DashboardPanelTimeseries{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "timeseries",
		},
		PluginVersion: "11.5.1",
		FieldConfig: schema.DashboardPanelFieldConfig[schema.DashboardPanelTimeseriesCustom]{
			Overrides: make([]schema.DashboardPanelFieldConfigOverride, 0),
			Defaults: schema.DashboardPanelFieldConfigDefaults[schema.DashboardPanelTimeseriesCustom]{
				Custom: &schema.DashboardPanelTimeseriesCustom{
					DrawStyle:         schema.GraphDrawStyleLine,
					LineInterpolation: schema.LineInterpolationLinear,
					BarAlignment:      schema.BarAlignmentCenter,
					LineWidth:         1.0,
					FillOpacity:       100.0,
					GradientMode:      schema.GraphGradientModeNone,
					ShowPoints:        schema.VisibilityModeAuto,
					PointSize:         1.0,
					AxisPlacement:     schema.AxisPlacementAuto,
					AxisColorMode:     schema.AxisColorModeText,
					ScaleDistribution: schema.ScaleDistributionConfig{
						Type: schema.ScaleDistributionLinear,
					},
					ThresholdsStyle: schema.GraphThresholdsStyleConfig{
						Mode: schema.GraphThresholdsStyleModeOff,
					},
					Stacking: schema.StackingConfig{
						Mode:  schema.StackingModeNone,
						Group: "A",
					},
				},
				Mappings: make([]schema.DashboardPanelFieldConfigMapping, 0),
			},
		},
		Options: schema.DashboardPanelTimeseriesOptions{
			Timezone: []string{"utc"},
			Legend: schema.DashboardPanelTimeseriesOptionsLegend{
				DisplayMode: schema.TimeseriesLegendDisplayModeList,
				Placement:   schema.TimeseriesLegendPlacementBottom,
				ShowLegend:  true,
			},
			Tooltip: schema.DashboardPanelTimeseriesOptionsTooltip{
				Mode: schema.TimeseriesTooltipModeSingle,
				Sort: schema.TimeseriesTooltipSortNone,
			},
		},
	}

	return &DashboardPanelTimeseriesBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		panel:                 &panel,
		nextRefID:             'A',
	}
}

func (b *DashboardPanelTimeseriesBuilder) Options(f func(*DashboardPanelTimeseriesOptionsBuilder)) {
	f(&DashboardPanelTimeseriesOptionsBuilder{parent: b})
}

func (b *DashboardPanelTimeseriesBuilder) Datasource(f func(*DashboardDatasourceRefBuilder)) {
	f(&DashboardDatasourceRefBuilder{parent: &b.panel.Datasource})
}

func (b *DashboardPanelTimeseriesBuilder) FieldConfig(f func(builder *DashboardPanelTimeseriesFieldConfigBuilder)) {
	f(&DashboardPanelTimeseriesFieldConfigBuilder{fieldConfig: &b.panel.FieldConfig})
}

func (b *DashboardPanelTimeseriesBuilder) Targets(fn func(builder *DashboardPanelTargetsBuilder)) {
	fn(&DashboardPanelTargetsBuilder{
		nextRefID: &b.nextRefID,
		targets:   &b.panel.Targets,
	})
}

func (b *DashboardPanelTargetsBuilder) AthenaSQL(datasourceUID string, fn func(builder *DashboardPanelTargetAthenaSQLBuilder)) {
	refID := string([]rune{*b.nextRefID})
	*b.nextRefID += 1

	var target schema.DashboardPanelGenericTargetImpl[schema.DashboardPanelTargetAthenaSQL]
	target.Data.RefID = refID
	target.Data.Format = schema.QueryFormatAthenaSQLTimeSeries
	target.Data.Datasource.Type = "grafana-athena-datasource"
	target.Data.Datasource.UID = datasourceUID
	target.Data.ConnectionArgs = schema.DashboardPanelTargetAthenaSQLConnectionArgs{
		Catalog:                    "__default",
		Database:                   "__default",
		Region:                     "__default",
		ResultReuseEnabled:         true,
		ResultReuseMaxAgeInMinutes: 60.0,
	}

	fn(&DashboardPanelTargetAthenaSQLBuilder{
		dashboardPanelTargetQueryBuilder: dashboardPanelTargetQueryBuilder{
			customJSON: &target.CustomJSON,
			target:     &target.Data.DashboardPanelTargetQuery,
		},
		target: &target.Data,
	})
	*b.targets = append(*b.targets, &target)
}

func (b *DashboardPanelTimeseriesBuilder) JSON(customJSON interface{}) {
	b.panel.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTimeseriesOptionsBuilder) Tooltip(f func(*DashboardPanelTimeseriesOptionsTooltipBuilder)) {
	f(&DashboardPanelTimeseriesOptionsTooltipBuilder{parent: b.parent})
}

func (b *DashboardPanelTimeseriesOptionsBuilder) Legend(f func(*DashboardPanelTimeseriesOptionsLegendBuilder)) {
	f(&DashboardPanelTimeseriesOptionsLegendBuilder{parent: b.parent})
}

func (b *DashboardPanelTimeseriesOptionsBuilder) JSON(customJSON interface{}) {
	b.parent.panel.Options.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTimeseriesOptionsTooltipBuilder) HideZeros(hideZeros bool) {
	b.parent.panel.Options.Tooltip.HideZeros = hideZeros
}

func (b *DashboardPanelTimeseriesOptionsTooltipBuilder) Mode(mode schema.TimeseriesTooltipMode) {
	b.parent.panel.Options.Tooltip.Mode = mode
}

func (b *DashboardPanelTimeseriesOptionsTooltipBuilder) Sort(sort schema.TimeseriesTooltipSort) {
	b.parent.panel.Options.Tooltip.Sort = sort
}

func (b *DashboardPanelTimeseriesOptionsTooltipBuilder) JSON(customJSON interface{}) {
	b.parent.panel.Options.Tooltip.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTimeseriesOptionsLegendBuilder) DisplayMode(displayMode schema.TimeseriesLegendDisplayMode) {
	b.parent.panel.Options.Legend.DisplayMode = displayMode
}

func (b *DashboardPanelTimeseriesOptionsLegendBuilder) Placement(placement schema.TimeseriesLegendPlacement) {
	b.parent.panel.Options.Legend.Placement = placement
}

func (b *DashboardPanelTimeseriesOptionsLegendBuilder) ShowLegend(showLegend bool) {
	b.parent.panel.Options.Legend.ShowLegend = showLegend
}

func (b *DashboardPanelTimeseriesOptionsLegendBuilder) JSON(customJSON interface{}) {
	b.parent.panel.Options.Legend.CustomJSON.Add(customJSON)
}
