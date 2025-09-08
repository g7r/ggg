package schema

import "encoding/json"

type GraphTooltip int

const (
	// No shared crosshair or tooltip (default)
	GraphTooltipDefault GraphTooltip = 0
	// Shared crosshair
	GraphTooltipSharedCrosshair GraphTooltip = 1
	// Shared crosshair AND shared tooltip.
	GraphTooltipSharedCrosshairAndTooltip GraphTooltip = 2
)

type DashboardStyle string

const (
	DashboardStyleDark  DashboardStyle = "dark"
	DashboardStyleLight DashboardStyle = "light"
)

type LinkIcon string

const (
	LinkIconExternalLink LinkIcon = "external link"
	LinkIconDashboard    LinkIcon = "dashboard"
	LinkIconQuestion     LinkIcon = "question"
	LinkIconInfo         LinkIcon = "info"
	LinkIconBolt         LinkIcon = "bolt"
	LinkIconDoc          LinkIcon = "doc"
	LinkIconCloud        LinkIcon = "cloud"
)

type DashboardVariableHide int

const (
	DashboardVariableDontHide     DashboardVariableHide = 0
	DashboardVariableHideLabel    DashboardVariableHide = 1
	DashboardVariableHideVariable DashboardVariableHide = 2
)

type DashboardVariableRefresh int

const (
	DashboardVariableRefreshNever              DashboardVariableRefresh = 0
	DashboardVariableRefreshOnDashboardLoad    DashboardVariableRefresh = 1
	DashboardVariableRefreshOnTimeRangeChanged DashboardVariableRefresh = 2
)

type DashboardVariableSort int

const (
	DashboardVariableSortDisabled                        DashboardVariableSort = 0
	DashboardVariableSortAlphabeticalAsc                 DashboardVariableSort = 1
	DashboardVariableSortAlphabeticalDesc                DashboardVariableSort = 2
	DashboardVariableSortNumericalAsc                    DashboardVariableSort = 3
	DashboardVariableSortNumericalDesc                   DashboardVariableSort = 4
	DashboardVariableSortAlphabeticalCaseInsensitiveAsc  DashboardVariableSort = 5
	DashboardVariableSortAlphabeticalCaseInsensitiveDesc DashboardVariableSort = 6
)

type GraphDrawStyle string

const (
	GraphDrawStyleBars   GraphDrawStyle = "bars"
	GraphDrawStyleLine   GraphDrawStyle = "line"
	GraphDrawStylePoints GraphDrawStyle = "points"
)

type GraphGradientMode string

const (
	GraphGradientModeNone    GraphGradientMode = "none"
	GraphGradientModeHue     GraphGradientMode = "hue"
	GraphGradientModeOpacity GraphGradientMode = "opacity"
	GraphGradientModeScheme  GraphGradientMode = "scheme"
)

type LineInterpolation string

const (
	LineInterpolationLinear     LineInterpolation = "linear"
	LineInterpolationSmooth     LineInterpolation = "smooth"
	LineInterpolationStepAfter  LineInterpolation = "stepAfter"
	LineInterpolationStepBefore LineInterpolation = "stepBefore"
)

type BarAlignment int

const (
	BarAlignmentBefore BarAlignment = -1
	BarAlignmentCenter BarAlignment = 0
	BarAlignmentAfter  BarAlignment = 1
)

type VisibilityMode string

const (
	VisibilityModeAuto   VisibilityMode = "auto"
	VisibilityModeAlways VisibilityMode = "always"
	VisibilityModeNever  VisibilityMode = "never"
)

type AxisPlacement string

const (
	AxisPlacementAuto   AxisPlacement = "auto"
	AxisPlacementBottom AxisPlacement = "bottom"
	AxisPlacementLeft   AxisPlacement = "left"
	AxisPlacementRight  AxisPlacement = "right"
	AxisPlacementTop    AxisPlacement = "top"
	AxisPlacementHidden AxisPlacement = "hidden"
)

type AxisColorMode string

const (
	AxisColorModeSeries AxisColorMode = "series"
	AxisColorModeText   AxisColorMode = "text"
)

type ScaleDistribution string

const (
	ScaleDistributionLinear  ScaleDistribution = "linear"
	ScaleDistributionLog     ScaleDistribution = "log"
	ScaleDistributionOrdinal ScaleDistribution = "ordinal"
	ScaleDistributionSymLog  ScaleDistribution = "symlog"
)

type GraphThresholdsStyleMode string

const (
	GraphThresholdsStyleModeArea          GraphThresholdsStyleMode = "area"
	GraphThresholdsStyleModeDashed        GraphThresholdsStyleMode = "dashed"
	GraphThresholdsStyleModeDashedAndArea GraphThresholdsStyleMode = "dashed+area"
	GraphThresholdsStyleModeLine          GraphThresholdsStyleMode = "line"
	GraphThresholdsStyleModeLineAndArea   GraphThresholdsStyleMode = "line+area"
	GraphThresholdsStyleModeOff           GraphThresholdsStyleMode = "off"
	GraphThresholdsStyleModeSeries        GraphThresholdsStyleMode = "series"
)

type StackingMode string

const (
	StackingModeNone    StackingMode = "none"
	StackingModeNormal  StackingMode = "normal"
	StackingModePercent StackingMode = "percent"
)

var DefaultDashboardRefreshIntervals = []string{"5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"}

type Dashboard struct {
	CustomJSON customJSON `json:"-"`

	// Version of the JSON schema, incremented each time a Grafana update brings
	// changes to said schema.
	SchemaVersion int `json:"schemaVersion"`

	// Human-readable title.
	Title string `json:"title,omitempty"`
	// Human-readable dashboard description.
	Description string `json:"description,omitempty"`
	// Unique dashboard ID, used in URL.
	UID string `json:"uid,omitempty"`
	// Time zone, use user default if empty.
	Timezone string `json:"timezone"`
	// Allow dashboard to be edited in Grafana Web UI.
	Editable bool `json:"editable"`
	// Dashboard ID from https://grafana.com/grafana/dashboards/
	GnetID *int64 `json:"gnetId,omitempty"`
	// Mode of tooltip sharing between graph panels.
	GraphTooltip GraphTooltip `json:"graphTooltip"`
	// Auto-refresh period, should be like 5s, 1m, 3h. Empty string means that auto-refresh is disabled.
	Refresh string `json:"refresh"`
	// Tags associated with dashboard.
	Tags []string `json:"tags,omitempty"`
	// Theme of dashboard.
	Style DashboardStyle `json:"style"`
	// Time range for dashboard, e.g. last 6 hours, last 7 days, etc (default: last 6 hours)
	Time *DashboardTime `json:"time,omitempty"`
	// Timepicker metadata.
	Timepicker *DashboardTimepicker `json:"timepicker,omitempty"`

	Annotations DashboardAnnotationList `json:"annotations,omitempty"`
	Links       []DashboardLink         `json:"links,omitempty"`
	Templating  DashboardVariableList   `json:"templating,omitempty"`
	Panels      []DashboardPanel        `json:"panels"`
}

type DashboardTime struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type DashboardTimepicker struct {
	// Whether timepicker is collapsed or not.
	Collapse bool `json:"collapse"`
	// Whether timepicker is enabled or not.
	Enable bool `json:"enable"`
	// Whether timepicker is visible or not.
	Hidden bool `json:"hidden"`
	// Selectable intervals for auto-refresh (default: 5s, 10s, 30s, 1m, 5m, 15m, 30m, 1h, 2h, 1d).
	RefreshIntervals []string `json:"refresh_intervals"`
	// Specify for example 5m to ignore last 5 minutes
	NowDelay string `json:"nowDelay,omitempty"`
}

type ScaleDistributionConfig struct {
	LinearThreshold *float64          `json:"linearThreshold,omitempty"`
	Log             *float64          `json:"log,omitempty"`
	Type            ScaleDistribution `json:"type"`
}

type HideSeriesConfig struct {
	Legend  bool `json:"legend"`
	Tooltip bool `json:"tooltip"`
	Viz     bool `json:"viz"`
}

type GraphThresholdsStyleConfig struct {
	Mode GraphThresholdsStyleMode `json:"mode"`
}

type StackingConfig struct {
	Group string       `json:"group"`
	Mode  StackingMode `json:"mode"`
}

type DashboardLink interface {
	isDashboardLink()
}

type DashboardVariable interface {
	isDashboardVariable()
}

type DashboardPanel interface {
	isDashboardPanel()
}

type DashboardAnnotationList []*DashboardAnnotation
type DashboardVariableList []DashboardVariable

func (d *Dashboard) MarshalJSON() ([]byte, error) {
	type plainDashboard Dashboard
	return marshalResource((*plainDashboard)(d), d.CustomJSON)
}

func (d *DashboardAnnotationList) MarshalJSON() ([]byte, error) {
	var annotationList struct {
		List []*DashboardAnnotation `json:"list"`
	}

	annotationList.List = *d

	return json.Marshal(annotationList)
}

func (d *DashboardVariableList) MarshalJSON() ([]byte, error) {
	var variableList struct {
		List []DashboardVariable `json:"list"`
	}

	variableList.List = *d

	return json.Marshal(variableList)
}
