package schema

type DashboardPanelGauge struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string `json:"pluginVersion"`
	Datasource    string `json:"datasource,omitempty"`

	// TODO: support field config
	FieldConfig DashboardPanelFieldConfig `json:"fieldConfig"`

	// Min interval. Recommended to be set to write frequency.
	Interval string `json:"interval,omitempty"`
	// Max data points
	MaxDataPoints int `json:"maxDataPoints,omitempty"`
	// Relative time
	TimeFrom string `json:"timeFrom,omitempty"`
	// Time shift
	TimeShift string `json:"timeShift,omitempty"`

	Targets []*DashboardPanelTarget `json:"targets"`
}

func (d *DashboardPanelGauge) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelGauge
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}
