package schema

type DashboardPanelTable struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string                  `json:"pluginVersion"`
	Targets       []*DashboardPanelTarget `json:"targets"`

	// TODO: support field config
	FieldConfig DashboardPanelFieldConfig `json:"fieldConfig"`
}

func (d *DashboardPanelTable) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTable
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}
