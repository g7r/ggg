package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTableBuilder struct {
	dashboardPanelBuilder
	panel *schema.DashboardPanelTable

	nextRefID rune
}

type DashboardPanelTableTargetsBuilder struct {
	parent *DashboardPanelTableBuilder
}

func newDashboardPanelTableBuilder() *DashboardPanelTableBuilder {
	panel := schema.DashboardPanelTable{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "table",
		},
		PluginVersion: "7.5.3",
	}

	return &DashboardPanelTableBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		panel:                 &panel,
		nextRefID:             'A',
	}
}

func (b *DashboardPanelTableBuilder) FieldConfig(fn func(*DashboardPanelFieldConfigBuilder)) {
	fn(&DashboardPanelFieldConfigBuilder{fieldConfig: &b.panel.FieldConfig})
}

func (b *DashboardPanelTableBuilder) Targets(fn func(builder *DashboardPanelTableTargetsBuilder)) {
	fn(&DashboardPanelTableTargetsBuilder{parent: b})
}

func (b *DashboardPanelTableBuilder) JSON(customJSON interface{}) {
	b.panel.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTableTargetsBuilder) AddTarget(fn func(builder *DashboardPanelTargetBuilder)) {
	refID := string([]rune{b.parent.nextRefID})
	b.parent.nextRefID += 1

	target := schema.DashboardPanelTarget{
		RefID:          refID,
		Format:         schema.TargetFormatTable,
		IntervalFactor: 1,
	}

	fn(&DashboardPanelTargetBuilder{target: &target})
	b.parent.panel.Targets = append(b.parent.panel.Targets, &target)
}
