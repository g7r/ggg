package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelRowBuilder struct {
	dashboardPanelBuilder
	parent *DashboardBuilder
	panel  *schema.DashboardPanelRow
}

func newDashboardPanelRowBuilder(parent *DashboardBuilder) *DashboardPanelRowBuilder {
	panel := schema.DashboardPanelRow{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "row",
		},
	}

	return &DashboardPanelRowBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		parent:                parent,
		panel:                 &panel,
	}
}

func (b *DashboardPanelRowBuilder) Collapsed(collapsed bool) {
	b.panel.Collapsed = collapsed
}

func (b *DashboardPanelRowBuilder) Title(title string) {
	b.panel.Title = title
}

func (b *DashboardPanelRowBuilder) Panels(fn func(*DashboardPanelsBuilder)) {
	fn(&DashboardPanelsBuilder{parent: b.parent, panels: &b.panel.Panels})
}
