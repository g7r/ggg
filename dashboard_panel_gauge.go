package ggg

import (
	"fmt"

	"github.com/g7r/ggg/schema"
)

type DashboardPanelGaugeBuilder struct {
	dashboardPanelBuilder
	panel *schema.DashboardPanelGauge

	nextRefID rune
}

type DashboardPanelGaugeTargetsBuilder struct {
	parent *DashboardPanelGaugeBuilder
}

func newDashboardPanelGaugeBuilder() *DashboardPanelGaugeBuilder {
	panel := schema.DashboardPanelGauge{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "gauge",
		},
		PluginVersion: "7.4.5",
	}

	return &DashboardPanelGaugeBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		panel:                 &panel,
		nextRefID:             'A',
	}
}

func (b *DashboardPanelGaugeBuilder) Datasource(datasource string) {
	b.panel.Datasource = datasource
}

func (b *DashboardPanelGaugeBuilder) DatasourceVar(datasourceVar string) {
	b.panel.Datasource = fmt.Sprintf("${%s}", datasourceVar)
}

func (b *DashboardPanelGaugeBuilder) MaxDataPoints(maxDataPoints int) {
	b.panel.MaxDataPoints = maxDataPoints
}

func (b *DashboardPanelGaugeBuilder) FieldConfig(fn func(*DashboardPanelFieldConfigBuilder)) {
	fn(&DashboardPanelFieldConfigBuilder{fieldConfig: &b.panel.FieldConfig})
}

func (b *DashboardPanelGaugeBuilder) Targets(fn func(builder *DashboardPanelGaugeTargetsBuilder)) {
	fn(&DashboardPanelGaugeTargetsBuilder{parent: b})
}

func (b *DashboardPanelGaugeTargetsBuilder) AddTarget(fn func(builder *DashboardPanelTargetBuilder)) {
	refID := string([]rune{b.parent.nextRefID})
	b.parent.nextRefID += 1

	target := schema.DashboardPanelTarget{
		RefID:          refID,
		Format:         schema.TargetFormatTimeSeries,
		IntervalFactor: 1,
	}

	fn(&DashboardPanelTargetBuilder{target: &target})
	b.parent.panel.Targets = append(b.parent.panel.Targets, &target)
}

func (b *DashboardPanelGaugeBuilder) JSON(customJSON interface{}) {
	b.panel.CustomJSON.Add(customJSON)
}
