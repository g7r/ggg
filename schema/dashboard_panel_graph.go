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

type Unit string

// TODO: add other predefined formats
const (
	UnitTimeNanos   Unit = "ns"
	UnitTimeMicros  Unit = "µs"
	UnitTimeMillis  Unit = "ms"
	UnitTimeSeconds Unit = "s"
	UnitTimeMinutes Unit = "m"
	UnitTimeHours   Unit = "h"
	UnitTimeDays    Unit = "d"

	// All the ones with the "Duration" prefix print a friendly duration name instead of small unit.
	UnitTimeDurationMillis  Unit = "dtdurationms"
	UnitTimeDurationSeconds Unit = "dtdurations"
	UnitTimeDurationHHMMSS  Unit = "dthms"
	UnitTimeDurationDHHMMSS Unit = "dtdhms"

	UnitTimeHertz     Unit = "hertz"
	UnitTimeTimeticks Unit = "timeticks"

	// Those with "Clock" prefix show durations as "01h:02m:03s:004ms"
	UnitTimeClockMillis  Unit = "clockms"
	UnitTimeClockSeconds Unit = "clocks"

	UnitThroughputCountsPerSecond   Unit = "cps"
	UnitThroughputOpsPerSecond      Unit = "ops"
	UnitThroughputRequestsPerSecond Unit = "reqps"
	UnitThroughputReadsPerSecond    Unit = "rps"
	UnitThroughputWritesPerSecond   Unit = "wps"
	UnitThroughputIOPerSecond       Unit = "iops"
	UnitThroughputCountsPerMinute   Unit = "cpm"
	UnitThroughputOpsPerMinute      Unit = "opm"
	UnitThroughputReadsPerMinute    Unit = "rpm"
	UnitThroughputWritesPerMinute   Unit = "wpm"

	UnitCurrencyDollars Unit = "currencyUSD"
	UnitCurrencyEuro    Unit = "currencyEUR"
	UnitCurrencyRubles  Unit = "currencyRUB"

	UnitDataRatePacketsPerSecond   Unit = "pps"
	UnitDataRateBytesPerSecondIEC  Unit = "binBps"
	UnitDataRateBytesPerSecondSI   Unit = "Bps"
	UnitDataRateBitsPerSecondIEC   Unit = "binbps"
	UnitDataRateBitsPerSecondSI    Unit = "bps"
	UnitDataRateKibibytesPerSecond Unit = "KiBs"
	UnitDataRateKibibitsPerSecond  Unit = "Kibits"
	UnitDataRateKilobytesPerSecond Unit = "KBs"
	UnitDataRateKilobitsPerSecond  Unit = "Kbits"
	UnitDataRateMebibytesPerSecond Unit = "MiBs"
	UnitDataRateMebibitsPerSecond  Unit = "Mibits"
	UnitDataRateMegabytesPerSecond Unit = "MBs"
	UnitDataRateMegabitsPerSecond  Unit = "Mbits"
	UnitDataRateGibibytesPerSecond Unit = "GiBs"
	UnitDataRateGibibitsPerSecond  Unit = "Gibits"
	UnitDataRateGigabytesPerSecond Unit = "GBs"
	UnitDataRateGigabitsPerSecond  Unit = "Gbits"
	UnitDataRateTebibytesPerSecond Unit = "TiBs"
	UnitDataRateTebibitsPerSecond  Unit = "Tibits"
	UnitDataRateTerabytesPerSecond Unit = "TBs"
	UnitDataRateTerabitsPerSecond  Unit = "Tbits"
	UnitDataRatePebibytesPerSecond Unit = "PiBs"
	UnitDataRatePebibitsPerSecond  Unit = "Pibits"
	UnitDataRatePetabytesPerSecond Unit = "PBs"
	UnitDataRatePetabitsPerSecond  Unit = "Pbits"

	UnitDataBytesIEC  Unit = "bytes"
	UnitDataBytesSI   Unit = "decbytes"
	UnitDataBitsIEC   Unit = "bits"
	UnitDataBitsSI    Unit = "decbits"
	UnitDataKibibytes Unit = "kbytes"
	UnitDataKilobytes Unit = "deckbytes"
	UnitDataMebibytes Unit = "mbytes"
	UnitDataMegabytes Unit = "decmbytes"
	UnitDataGibibytes Unit = "gbytes"
	UnitDataGigabytes Unit = "decgbytes"
	UnitDataTebibytes Unit = "tbytes"
	UnitDataTerabytes Unit = "dectbytes"
	UnitDataPebibytes Unit = "pbytes"
	UnitDataPetabytes Unit = "decpbytes"

	UnitMiscNone               Unit = "none"
	UnitMiscString             Unit = "string"
	UnitMiscShort              Unit = "short"
	UnitMiscPercent100         Unit = "percent"
	UnitMiscPercent1           Unit = "percentunit"
	UnitMiscHexadecimal0x      Unit = "hex0x"
	UnitMiscHexadecimal        Unit = "hex"
	UnitMiscScientificNotation Unit = "sci"
	UnitMiscLocaleFormat       Unit = "locale"
	UnitMiscPixels             Unit = "pixel"
	UnitMiscDecibels           Unit = "dB"
)

type DashboardPanelGraph struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string `json:"pluginVersion"`
	Renderer      string `json:"renderer"`
	Datasource    string `json:"datasource,omitempty"`

	Bars          bool          `json:"bars,omitempty"`
	Lines         bool          `json:"lines"`
	Dashes        bool          `json:"dashes,omitempty"`
	Points        bool          `json:"points,omitempty"`
	SteppedLine   bool          `json:"steppedLine,omitempty"`
	Stack         bool          `json:"stack,omitempty"`
	Percentage    bool          `json:"percentage,omitempty"`
	NullPointMode NullPointMode `json:"nullPointMode"`
	Decimals      *int          `json:"decimals,omitempty"`

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
	Targets      []*DashboardPanelTarget         `json:"targets"`

	// TODO: support time regions
	TimeRegions []interface{} `json:"timeRegions,omitempty"`

	// TODO: support thresholds
	Thresholds []interface{} `json:"thresholds,omitempty"`

	// TODO: support series overrides
	SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`

	// TODO: support links
	Links []interface{} `json:"links,omitempty"`

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

	Sort     string `json:"sort,omitempty"`
	SortDesc bool   `json:"sortDesc,omitempty"`
}

type DashboardPanelGraphYAxes struct {
	Left  DashboardPanelGraphYAxis `json:"-"`
	Right DashboardPanelGraphYAxis `json:"-"`
}

type DashboardPanelGraphYAxis struct {
	CustomJSON customJSON `json:"-"`

	Show    bool     `json:"show"`
	Format  Unit     `json:"format"`
	Label   string   `json:"label,omitempty"`
	LogBase float64  `json:"logBase"`
	Min     *float64 `json:"min,omitempty"`
	Max     *float64 `json:"max,omitempty"`
	// Override automatic decimal precision
	Decimals *int `json:"decimals,omitempty"`
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

func (d *DashboardPanelGraphYAxes) MarshalJSON() ([]byte, error) {
	return json.Marshal([2]*DashboardPanelGraphYAxis{&d.Left, &d.Right})
}

func (d *DashboardPanelGraphYAxis) MarshalJSON() ([]byte, error) {
	type plainYAxis DashboardPanelGraphYAxis
	return marshalResource((*plainYAxis)(d), d.CustomJSON)
}
