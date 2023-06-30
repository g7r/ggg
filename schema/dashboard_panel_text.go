package schema

type TextMode string

const (
	TextModeMarkdown TextMode = "markdown"
	TextModeHTML     TextMode = "html"
)

type DashboardPanelText struct {
	CustomJSON customJSON `json:"-"`

	DashboardPanel             `json:"-"`
	DashboardPanelCommonFields `json:",inline"`

	PluginVersion string                    `json:"pluginVersion"`
	Options       DashboardPanelTextOptions `json:"options"`
}

type DashboardPanelTextOptions struct {
	CustomJSON customJSON `json:"-"`

	Content string   `json:"content"`
	Mode    TextMode `json:"mode"`
}

func (d *DashboardPanelText) MarshalJSON() ([]byte, error) {
	type plainPanel DashboardPanelText
	return marshalResource((*plainPanel)(d), d.CustomJSON)
}

func (d *DashboardPanelTextOptions) MarshalJSON() ([]byte, error) {
	type plainOptions DashboardPanelTextOptions
	return marshalResource((*plainOptions)(d), d.CustomJSON)
}
