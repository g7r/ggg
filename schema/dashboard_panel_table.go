package schema

type DashboardPanelTable struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string                  `json:"pluginVersion"`
	Targets       []*DashboardPanelTarget `json:"targets"`

	// TODO: support field config
	FieldConfig DashboardPanelFieldConfig[DashboardPanelTableCustom] `json:"fieldConfig"`
}

type DashboardPanelTableCustom struct {
	CustomJSON customJSON `json:"-"`
}

func (d *DashboardPanelTable) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelTable
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTableCustom) MarshalJSON() ([]byte, error) {
	type plainCustom DashboardPanelTableCustom
	return marshalResource((*plainCustom)(d), d.CustomJSON)
}
