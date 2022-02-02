package schema

import "encoding/json"

type NullPointMode string

const (
	NullPointModeNull       NullPointMode = "null"
	NullPointModeConnected  NullPointMode = "connected"
	NullPointModeNullAsZero NullPointMode = "null as zero"
)

type XAxisValue string

const (
	XAxisValueTotal   XAxisValue = "total"
	XAxisValueAvg     XAxisValue = "avg"
	XAxisValueMin     XAxisValue = "min"
	XAxisValueMax     XAxisValue = "max"
	XAxisValueCount   XAxisValue = "count"
	XAxisValueCurrent XAxisValue = "current"
)

type GraphTooltipSort int

const (
	GraphTooltipSortNone       GraphTooltipSort = 0
	GraphTooltipSortIncreasing GraphTooltipSort = 1
	GraphTooltipSortDecreasing GraphTooltipSort = 2
)

type GraphYAxisFormat string

// TODO: add other predefined formats
const (
	GraphYAxisFormatTimeNanos   GraphYAxisFormat = "ns"
	GraphYAxisFormatTimeMicros  GraphYAxisFormat = "µs"
	GraphYAxisFormatTimeMillis  GraphYAxisFormat = "ms"
	GraphYAxisFormatTimeSeconds GraphYAxisFormat = "s"
	GraphYAxisFormatTimeMinutes GraphYAxisFormat = "m"
	GraphYAxisFormatTimeHours   GraphYAxisFormat = "h"
	GraphYAxisFormatTimeDays    GraphYAxisFormat = "d"

	// All the ones with the "Duration" prefix print a friendly duration name instead of small unit.
	GraphYAxisFormatTimeDurationMillis  GraphYAxisFormat = "dtdurationms"
	GraphYAxisFormatTimeDurationSeconds GraphYAxisFormat = "dtdurations"
	GraphYAxisFormatTimeDurationHHMMSS  GraphYAxisFormat = "dthms"
	GraphYAxisFormatTimeDurationDHHMMSS GraphYAxisFormat = "dtdhms"

	GraphYAxisFormatTimeHertz     GraphYAxisFormat = "hertz"
	GraphYAxisFormatTimeTimeticks GraphYAxisFormat = "timeticks"

	// Those with "Clock" prefix show durations as "01h:02m:03s:004ms"
	GraphYAxisFormatTimeClockMillis  GraphYAxisFormat = "clockms"
	GraphYAxisFormatTimeClockSeconds GraphYAxisFormat = "clocks"

	GraphYAxisFormatThroughputCountsPerSecond   GraphYAxisFormat = "cps"
	GraphYAxisFormatThroughputOpsPerSecond      GraphYAxisFormat = "ops"
	GraphYAxisFormatThroughputRequestsPerSecond GraphYAxisFormat = "reqps"
	GraphYAxisFormatThroughputReadsPerSecond    GraphYAxisFormat = "rps"
	GraphYAxisFormatThroughputWritesPerSecond   GraphYAxisFormat = "wps"
	GraphYAxisFormatThroughputIOPerSecond       GraphYAxisFormat = "iops"
	GraphYAxisFormatThroughputCountsPerMinute   GraphYAxisFormat = "cpm"
	GraphYAxisFormatThroughputOpsPerMinute      GraphYAxisFormat = "opm"
	GraphYAxisFormatThroughputReadsPerMinute    GraphYAxisFormat = "rpm"
	GraphYAxisFormatThroughputWritesPerMinute   GraphYAxisFormat = "wpm"

	GraphYAxisFormatCurrencyDollars GraphYAxisFormat = "currencyUSD"
	GraphYAxisFormatCurrencyEuro    GraphYAxisFormat = "currencyEUR"
	GraphYAxisFormatCurrencyRubles  GraphYAxisFormat = "currencyRUB"

	GraphYAxisFormatDataRatePacketsPerSecond   GraphYAxisFormat = "pps"
	GraphYAxisFormatDataRateBytesPerSecondIEC  GraphYAxisFormat = "binBps"
	GraphYAxisFormatDataRateBytesPerSecondSI   GraphYAxisFormat = "Bps"
	GraphYAxisFormatDataRateBitsPerSecondIEC   GraphYAxisFormat = "binbps"
	GraphYAxisFormatDataRateBitsPerSecondSI    GraphYAxisFormat = "bps"
	GraphYAxisFormatDataRateKibibytesPerSecond GraphYAxisFormat = "KiBs"
	GraphYAxisFormatDataRateKibibitsPerSecond  GraphYAxisFormat = "Kibits"
	GraphYAxisFormatDataRateKilobytesPerSecond GraphYAxisFormat = "KBs"
	GraphYAxisFormatDataRateKilobitsPerSecond  GraphYAxisFormat = "Kbits"
	GraphYAxisFormatDataRateMebibytesPerSecond GraphYAxisFormat = "MiBs"
	GraphYAxisFormatDataRateMebibitsPerSecond  GraphYAxisFormat = "Mibits"
	GraphYAxisFormatDataRateMegabytesPerSecond GraphYAxisFormat = "MBs"
	GraphYAxisFormatDataRateMegabitsPerSecond  GraphYAxisFormat = "Mbits"
	GraphYAxisFormatDataRateGibibytesPerSecond GraphYAxisFormat = "GiBs"
	GraphYAxisFormatDataRateGibibitsPerSecond  GraphYAxisFormat = "Gibits"
	GraphYAxisFormatDataRateGigabytesPerSecond GraphYAxisFormat = "GBs"
	GraphYAxisFormatDataRateGigabitsPerSecond  GraphYAxisFormat = "Gbits"
	GraphYAxisFormatDataRateTebibytesPerSecond GraphYAxisFormat = "TiBs"
	GraphYAxisFormatDataRateTebibitsPerSecond  GraphYAxisFormat = "Tibits"
	GraphYAxisFormatDataRateTerabytesPerSecond GraphYAxisFormat = "TBs"
	GraphYAxisFormatDataRateTerabitsPerSecond  GraphYAxisFormat = "Tbits"
	GraphYAxisFormatDataRatePebibytesPerSecond GraphYAxisFormat = "PiBs"
	GraphYAxisFormatDataRatePebibitsPerSecond  GraphYAxisFormat = "Pibits"
	GraphYAxisFormatDataRatePetabytesPerSecond GraphYAxisFormat = "PBs"
	GraphYAxisFormatDataRatePetabitsPerSecond  GraphYAxisFormat = "Pbits"

	GraphYAxisFormatDataBytesIEC  GraphYAxisFormat = "bytes"
	GraphYAxisFormatDataBytesSI   GraphYAxisFormat = "decbytes"
	GraphYAxisFormatDataBitsIEC   GraphYAxisFormat = "bits"
	GraphYAxisFormatDataBitsSI    GraphYAxisFormat = "decbits"
	GraphYAxisFormatDataKibibytes GraphYAxisFormat = "kbytes"
	GraphYAxisFormatDataKilobytes GraphYAxisFormat = "deckbytes"
	GraphYAxisFormatDataMebibytes GraphYAxisFormat = "mbytes"
	GraphYAxisFormatDataMegabytes GraphYAxisFormat = "decmbytes"
	GraphYAxisFormatDataGibibytes GraphYAxisFormat = "gbytes"
	GraphYAxisFormatDataGigabytes GraphYAxisFormat = "decgbytes"
	GraphYAxisFormatDataTebibytes GraphYAxisFormat = "tbytes"
	GraphYAxisFormatDataTerabytes GraphYAxisFormat = "dectbytes"
	GraphYAxisFormatDataPebibytes GraphYAxisFormat = "pbytes"
	GraphYAxisFormatDataPetabytes GraphYAxisFormat = "decpbytes"

	GraphYAxisFormatMiscNone               GraphYAxisFormat = "none"
	GraphYAxisFormatMiscString             GraphYAxisFormat = "string"
	GraphYAxisFormatMiscShort              GraphYAxisFormat = "short"
	GraphYAxisFormatMiscPercent100         GraphYAxisFormat = "percent"
	GraphYAxisFormatMiscPercent1           GraphYAxisFormat = "percentunit"
	GraphYAxisFormatMiscHexadecimal0x      GraphYAxisFormat = "hex0x"
	GraphYAxisFormatMiscHexadecimal        GraphYAxisFormat = "hex"
	GraphYAxisFormatMiscScientificNotation GraphYAxisFormat = "sci"
	GraphYAxisFormatMiscLocaleFormat       GraphYAxisFormat = "locale"
	GraphYAxisFormatMiscPixels             GraphYAxisFormat = "pixel"
	GraphYAxisFormatMiscDecibels           GraphYAxisFormat = "dB"
)

type GraphTargetFormat string

const (
	// TODO: support other graph target formats
	GraphTargetFormatTimeSeries GraphTargetFormat = "time_series"
)

type DashboardPanelCommonFields struct {
	Type string `json:"type"`

	ID    int    `json:"id"`
	Title string `json:"title"`

	GridPos DashboardPanelGridPos `json:"gridPos"`
}

type DashboardPanelGridPos struct {
	CustomJSON customJSON `json:"-"`

	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
}

type DashboardPanelGraph struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string `json:"pluginVersion"`
	Renderer      string `json:"renderer"`
	Datasource    string `json:"datasource,omitempty"`

	Bars          bool          `json:"bars,omitempty"`
	Lines         bool          `json:"lines,omitempty"`
	Dashes        bool          `json:"dashes,omitempty"`
	Points        bool          `json:"points,omitempty"`
	SteppedLine   bool          `json:"steppedLine,omitempty"`
	Stack         bool          `json:"stack,omitempty"`
	Percentage    bool          `json:"percentage,omitempty"`
	NullPointMode NullPointMode `json:"nullPointMode"`

	// Area fill, only when Lines is true [0; 10], specifies fill opacity
	Fill int `json:"fill"`
	// Fill gradient, only when Fill>0 [0; 10]
	FillGradient int `json:"fillGradient"`
	// Line width, only when Lines is true [0; 10]
	LineWidth int `json:"linewidth"`
	// Point radius, only when Points is true [0.5; 10]
	PointRadius float64 `json:"pointradius"`

	DashLength  int `json:"dashLength,omitempty"`
	SpaceLength int `json:"spaceLength,omitempty"`

	HiddenSeries bool `json:"hiddenSeries"`

	XAxis DashboardPanelGraphXAxis `json:"xaxis"`

	// Min interval. Recommended to be set to write frequency.
	Interval string `json:"interval,omitempty"`
	// Max data points
	MaxDataPoints int `json:"maxDataPoints,omitempty"`
	// Relative time
	TimeFrom string `json:"timeFrom,omitempty"`
	// Time shift
	TimeShift string `json:"timeShift,omitempty"`

	Tooltip      DashboardPanelGraphTooltip      `json:"tooltip"`
	YAxisOptions DashboardPanelGraphYAxisOptions `json:"yaxis"`
	Options      DashboardPanelGraphOptions      `json:"options"`
	Legend       DashboardPanelGraphLegend       `json:"legend"`
	YAxes        DashboardPanelGraphYAxes        `json:"yaxes"`
	Targets      []*DashboardPanelGraphTarget    `json:"targets"`

	// TODO: support time regions
	TimeRegions []interface{} `json:"timeRegions,omitempty"`

	// TODO: support thresholds
	Thresholds []interface{} `json:"thresholds,omitempty"`

	// TODO: support series overrides
	SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`

	// TODO: support links
	Links []interface{} `json:"links,omitempty"`

	// TODO: support field config
	FieldConfig DashboardPanelGraphFieldConfig `json:"fieldConfig"`

	// TODO: support alias colors
	AliasColors struct{} `json:"aliasColors"`
}

type DashboardPanelGraphXAxis interface {
	isDashboardPanelGraphXAxis()
}

type DashboardPanelGraphXAxisCommonFields struct {
	Show bool   `json:"show"`
	Mode string `json:"mode"`
}

type DashboardPanelGraphXAxisTime struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanelGraphXAxis             `json:"-"`
	DashboardPanelGraphXAxisCommonFields `json:",inline"`
}

type DashboardPanelGraphXAxisSeries struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanelGraphXAxis             `json:"-"`
	DashboardPanelGraphXAxisCommonFields `json:",inline"`

	Values []XAxisValue `json:"values"`
}

type DashboardPanelGraphXAxisHistogram struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanelGraphXAxis             `json:"-"`
	DashboardPanelGraphXAxisCommonFields `json:",inline"`

	Buckets int     `json:"buckets,omitempty"`
	Min     float64 `json:"min,omitempty"`
	Max     float64 `json:"max,omitempty"`
}

type DashboardPanelGraphTooltip struct {
	CustomJSON customJSON `json:"-"`

	// Show all series on tooltip if true
	Shared bool `json:"shared"`
	// Tooltip values sorting order
	Sort GraphTooltipSort `json:"sort"`
	// UNKNOWN (always equal to "individual")
	ValueType string `json:"value_type"`
}

type DashboardPanelGraphYAxisOptions struct {
	CustomJSON customJSON `json:"-"`

	Align      bool    `json:"align"`
	AlignLevel float64 `json:"alignLevel,omitempty"`
}

type DashboardPanelGraphOptions struct {
	CustomJSON customJSON `json:"-"`

	AlertThreshold bool `json:"alertThreshold"`
}

type DashboardPanelGraphLegend struct {
	CustomJSON customJSON `json:"-"`

	Show bool `json:"show"`

	// Should be true if any other field below is true
	Values  bool `json:"values"`
	Avg     bool `json:"avg"`
	Current bool `json:"current"`
	Max     bool `json:"max"`
	Min     bool `json:"min"`
	Total   bool `json:"total"`

	AlignAsTable bool    `json:"alignAsTable,omitempty"`
	RightSide    bool    `json:"rightSide,omitempty"`
	SideWidth    float64 `json:"sideWidth,omitempty"`
}

type DashboardPanelGraphFieldConfig struct {
	CustomJSON customJSON `json:"-"`
}

type DashboardPanelGraphYAxes struct {
	Left  DashboardPanelGraphYAxis `json:"-"`
	Right DashboardPanelGraphYAxis `json:"-"`
}

type DashboardPanelGraphYAxis struct {
	CustomJSON customJSON `json:"-"`

	Show    bool             `json:"show"`
	Format  GraphYAxisFormat `json:"format"`
	Label   string           `json:"label,omitempty"`
	LogBase float64          `json:"logBase"`
	Min     *float64         `json:"min,omitempty"`
	Max     *float64         `json:"max,omitempty"`
	// Override automatic decimal precision
	Decimals int `json:"decimals,omitempty"`
}

type DashboardPanelGraphTarget struct {
	CustomJSON customJSON `json:"-"`

	RefID  string            `json:"refId"`
	Format GraphTargetFormat `json:"format"`
	Hide   bool              `json:"hide,omitempty"`
	Expr   string            `json:"expr"`
	// Min step (should be like "5s", "15m")
	Interval string `json:"interval,omitempty"`
	// Resolution is 1/IntervalFactor
	IntervalFactor int    `json:"intervalFactor"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Exemplar       bool   `json:"exemplar,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
}

func (d *DashboardPanelGridPos) MarshalJSON() ([]byte, error) {
	type plainGridPos DashboardPanelGridPos
	return marshalResource((*plainGridPos)(d), d.CustomJSON)
}

func (d *DashboardPanelGraph) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelGraph
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphXAxisTime) MarshalJSON() ([]byte, error) {
	type plainAxis DashboardPanelGraphXAxisTime
	return marshalResource((*plainAxis)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphXAxisSeries) MarshalJSON() ([]byte, error) {
	type plainAxis DashboardPanelGraphXAxisSeries
	return marshalResource((*plainAxis)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphXAxisHistogram) MarshalJSON() ([]byte, error) {
	type plainAxis DashboardPanelGraphXAxisHistogram
	return marshalResource((*plainAxis)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphTooltip) MarshalJSON() ([]byte, error) {
	type plainTooltip DashboardPanelGraphTooltip
	return marshalResource((*plainTooltip)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphYAxisOptions) MarshalJSON() ([]byte, error) {
	type plainOptions DashboardPanelGraphYAxisOptions
	return marshalResource((*plainOptions)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphOptions) MarshalJSON() ([]byte, error) {
	type plainOptions DashboardPanelGraphOptions
	return marshalResource((*plainOptions)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphLegend) MarshalJSON() ([]byte, error) {
	type plainLegend DashboardPanelGraphLegend
	return marshalResource((*plainLegend)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphFieldConfig) MarshalJSON() ([]byte, error) {
	type plainConfig DashboardPanelGraphFieldConfig
	return marshalResource((*plainConfig)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphYAxes) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]*DashboardPanelGraphYAxis{&d.Left, &d.Right})
}

func (d *DashboardPanelGraphYAxis) MarshalJSON() ([]byte, error) {
	type plainYAxis DashboardPanelGraphYAxis
	return marshalResource((*plainYAxis)(d), d.CustomJSON)
}

func (d *DashboardPanelGraphTarget) MarshalJSON() ([]byte, error) {
	type plainTarget DashboardPanelGraphTarget
	return marshalResource((*plainTarget)(d), d.CustomJSON)
}
