package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTextBuilder struct {
	dashboardPanelBuilder
	panel *schema.DashboardPanelText
}

type DashboardPanelTextOptionsBuilder struct {
	options *schema.DashboardPanelTextOptions
}

func newDashboardPanelTextBuilder() *DashboardPanelTextBuilder {
	panel := schema.DashboardPanelText{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "text",
		},
		PluginVersion: "7.4.5",
		Options:       schema.DashboardPanelTextOptions{Mode: schema.TextModeMarkdown},
	}

	return &DashboardPanelTextBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		panel:                 &panel,
	}
}

func (b *DashboardPanelTextBuilder) Options(fn func(*DashboardPanelTextOptionsBuilder)) {
	fn(&DashboardPanelTextOptionsBuilder{options: &b.panel.Options})
}

func (b *DashboardPanelTextBuilder) JSON(customJSON interface{}) {
	b.panel.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTextOptionsBuilder) Content(content string) {
	b.options.Content = content
}

func (b *DashboardPanelTextOptionsBuilder) Mode(mode schema.TextMode) {
	b.options.Mode = mode
}

func (b *DashboardPanelTextOptionsBuilder) JSON(customJSON interface{}) {
	b.options.CustomJSON.Add(customJSON)
}
