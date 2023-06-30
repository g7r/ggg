package schema

type TargetFormat string

const (
	// TODO: support other graph target formats
	TargetFormatTimeSeries TargetFormat = "time_series"
	TargetFormatTable      TargetFormat = "table"
)

type DashboardPanelTarget struct {
	CustomJSON customJSON `json:"-"`

	RefID  string       `json:"refId"`
	Format TargetFormat `json:"format"`
	Hide   bool         `json:"hide,omitempty"`
	Expr   string       `json:"expr,omitempty"`
	// Min step (should be like "5s", "15m")
	Interval string `json:"interval,omitempty"`
	// Resolution is 1/IntervalFactor
	IntervalFactor int    `json:"intervalFactor"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Exemplar       bool   `json:"exemplar,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
}

func (d *DashboardPanelTarget) MarshalJSON() ([]byte, error) {
	type plainTarget DashboardPanelTarget
	return marshalResource((*plainTarget)(d), d.CustomJSON)
}
