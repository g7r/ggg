package schema

type DashboardPanelRow struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	Collapsed bool             `json:"collapsed"`
	Panels    []DashboardPanel `json:"panels"`
}

func (d *DashboardPanelRow) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelRow
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}
