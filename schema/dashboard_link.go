package schema

type DashboardLinkCommonFields struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	IncludeVars bool   `json:"includeVars,omitempty"`
	KeepTime    bool   `json:"keepTime,omitempty"`
	TargetBlank bool   `json:"targetBlank,omitempty"`
}

type DashboardLinkDashboards struct {
	CustomJSON customJSON `json:"-"`

	DashboardLink             `json:"-"`
	DashboardLinkCommonFields `json:",inline"`

	AsDropdown bool     `json:"asDropdown"`
	Tags       []string `json:"tags"`
}

type DashboardLinkExternal struct {
	CustomJSON customJSON `json:"-"`

	DashboardLink             `json:"-"`
	DashboardLinkCommonFields `json:",inline"`

	URL     string   `json:"url"`
	Icon    LinkIcon `json:"icon"`
	Tooltip string   `json:"tooltip,omitempty"`
}

func (d *DashboardLinkDashboards) MarshalJSON() ([]byte, error) {
	type plainLink DashboardLinkDashboards
	return marshalResource((*plainLink)(d), d.CustomJSON)
}

func (d *DashboardLinkExternal) MarshalJSON() ([]byte, error) {
	type plainLink DashboardLinkExternal
	return marshalResource((*plainLink)(d), d.CustomJSON)
}
