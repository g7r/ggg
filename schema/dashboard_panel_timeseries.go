package schema

type TimeseriesTooltipMode string
type TimeseriesTooltipSort string
type TimeseriesLegendDisplayMode string
type TimeseriesLegendPlacement string

const (
	TimeseriesTooltipModeNone   TimeseriesTooltipMode = "none"
	TimeseriesTooltipModeSingle TimeseriesTooltipMode = "single"
	TimeseriesTooltipModeMulti  TimeseriesTooltipMode = "multi"
)

const (
	TimeseriesTooltipSortNone TimeseriesTooltipSort = "none"
	TimeseriesTooltipSortAsc  TimeseriesTooltipSort = "asc"
	TimeseriesTooltipSortDesc TimeseriesTooltipSort = "desc"
)

const (
	TimeseriesLegendDisplayModeList  TimeseriesLegendDisplayMode = "list"
	TimeseriesLegendDisplayModeTable TimeseriesLegendDisplayMode = "table"
)

const (
	TimeseriesLegendPlacementBottom TimeseriesLegendPlacement = "bottom"
	TimeseriesLegendPlacementRight  TimeseriesLegendPlacement = "right"
)

type DashboardPanelTimeseries struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string `json:"pluginVersion"`

	Options     DashboardPanelTimeseriesOptions                           `json:"options"`
	Datasource  DashboardDatasourceRef                                    `json:"datasource"`
	FieldConfig DashboardPanelFieldConfig[DashboardPanelTimeseriesCustom] `json:"fieldConfig"`
	Targets     []DashboardPanelGenericTarget                             `json:"targets"`
}

type DashboardPanelTimeseriesCustom struct {
	CustomJSON customJSON `json:"-"`

	DrawStyle         GraphDrawStyle             `json:"drawStyle"`
	LineInterpolation LineInterpolation          `json:"lineInterpolation"`
	BarAlignment      BarAlignment               `json:"barAlignment"`
	BarWidthFactor    *float64                   `json:"barWidthFactor,omitempty"`
	LineWidth         float64                    `json:"lineWidth"`
	FillOpacity       float64                    `json:"fillOpacity"`
	GradientMode      GraphGradientMode          `json:"gradientMode"`
	SpanNulls         bool                       `json:"spanNulls"`
	InsertNulls       bool                       `json:"insertNulls"`
	ShowPoints        VisibilityMode             `json:"showPoints"`
	PointSize         float64                    `json:"pointSize"`
	AxisPlacement     AxisPlacement              `json:"axisPlacement"`
	AxisLabel         string                     `json:"axisLabel"`
	AxisColorMode     AxisColorMode              `json:"axisColorMode"`
	AxisBorderShow    bool                       `json:"axisBorderShow"`
	AxisCenteredZero  bool                       `json:"axisCenteredZero"`
	ScaleDistribution ScaleDistributionConfig    `json:"scaleDistribution"`
	HideFrom          HideSeriesConfig           `json:"hideFrom"`
	ThresholdsStyle   GraphThresholdsStyleConfig `json:"thresholdsStyle"`
	Stacking          StackingConfig             `json:"stacking"`
}

type DashboardPanelTimeseriesOptions struct {
	CustomJSON customJSON `json:"-"`

	Legend   DashboardPanelTimeseriesOptionsLegend  `json:"legend"`
	Timezone []string                               `json:"timezone"`
	Tooltip  DashboardPanelTimeseriesOptionsTooltip `json:"tooltip"`
}

type DashboardPanelTimeseriesOptionsLegend struct {
	CustomJSON customJSON `json:"-"`

	DisplayMode TimeseriesLegendDisplayMode `json:"displayMode"`
	Placement   TimeseriesLegendPlacement   `json:"placement"`
	ShowLegend  bool                        `json:"showLegend"`
}

type DashboardPanelTimeseriesOptionsTooltip struct {
	CustomJSON customJSON `json:"-"`

	HideZeros bool                  `json:"hideZeros"`
	Mode      TimeseriesTooltipMode `json:"mode"`
	Sort      TimeseriesTooltipSort `json:"sort"`
}

func (d *DashboardPanelTimeseries) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTimeseries
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTimeseriesOptions) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTimeseriesOptions
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTimeseriesOptionsLegend) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTimeseriesOptionsLegend
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTimeseriesOptionsTooltip) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTimeseriesOptionsTooltip
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTimeseriesCustom) MarshalJSON() ([]byte, error) {
	type plainCustom DashboardPanelTimeseriesCustom
	return marshalResource((*plainCustom)(d), d.CustomJSON)
}
