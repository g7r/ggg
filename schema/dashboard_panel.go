package schema

type DashboardPanelCommonFields struct {
	Type string `json:"type"`

	ID    int    `json:"id"`
	Title string `json:"title"`

	GridPos     DashboardPanelGridPos `json:"gridPos"`
	Transparent bool                  `json:"transparent,omitempty"`
}

type DashboardPanelGridPos struct {
	CustomJSON customJSON `json:"-"`

	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"w"`
	Height int `json:"h"`
}

type DashboardPanelFieldConfigColor struct {
	CustomJSON customJSON `json:"-"`

	FixedColor string               `json:"fixedColor,omitempty"`
	Mode       FieldConfigColorMode `json:"mode"`
}

func (d *DashboardPanelGridPos) MarshalJSON() ([]byte, error) {
	type plainGridPos DashboardPanelGridPos
	return marshalResource((*plainGridPos)(d), d.CustomJSON)
}
