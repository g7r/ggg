package ggg

import (
	"crypto/sha256"

	"github.com/g7r/ggg/schema"
)

type DashboardBuilder struct {
	dashboard    *schema.Dashboard
	usedPanelIDs map[int]bool
}

type DashboardLinksBuilder struct {
	dashboard *schema.Dashboard
}

type DashboardAnnotationsBuilder struct {
	dashboard *schema.Dashboard
}

type DashboardTemplatingBuilder struct {
	dashboard *schema.Dashboard
}

type DashboardPanelsBuilder struct {
	parent *DashboardBuilder
	panels *[]schema.DashboardPanel
}

func NewDashboard(fn func(builder *DashboardBuilder)) *schema.Dashboard {
	builder := DashboardBuilder{
		dashboard: &schema.Dashboard{
			SchemaVersion: 40,
			Style:         schema.DashboardStyleLight,
		},
		usedPanelIDs: map[int]bool{},
	}

	fn(&builder)
	return builder.dashboard
}

func (b *DashboardBuilder) SchemaVersion(version int) {
	b.dashboard.SchemaVersion = version
}

func (b *DashboardBuilder) Title(title string) {
	b.dashboard.Title = title
}

func (b *DashboardBuilder) Description(description string) {
	b.dashboard.Description = description
}

func (b *DashboardBuilder) UID(uid string) {
	b.dashboard.UID = uid
}

func (b *DashboardBuilder) Editable(editable bool) {
	b.dashboard.Editable = editable
}

func (b *DashboardBuilder) GraphTooltip(tooltip schema.GraphTooltip) {
	b.dashboard.GraphTooltip = tooltip
}

func (b *DashboardBuilder) Refresh(refresh string) {
	b.dashboard.Refresh = refresh
}

func (b *DashboardBuilder) Time(from, to string) {
	b.dashboard.Time = &schema.DashboardTime{From: from, To: to}
}

func (b *DashboardBuilder) Timepicker(fn func(*DashboardTimepickerBuilder)) {
	if b.dashboard.Timepicker == nil {
		b.dashboard.Timepicker = &schema.DashboardTimepicker{
			Enable:           true,
			RefreshIntervals: schema.DefaultDashboardRefreshIntervals,
		}
	}

	fn(&DashboardTimepickerBuilder{timepicker: b.dashboard.Timepicker})
}

func (b *DashboardBuilder) Links(fn func(*DashboardLinksBuilder)) {
	fn(&DashboardLinksBuilder{dashboard: b.dashboard})
}

func (b *DashboardBuilder) Annotations(fn func(*DashboardAnnotationsBuilder)) {
	fn(&DashboardAnnotationsBuilder{dashboard: b.dashboard})
}

func (b *DashboardBuilder) Templating(fn func(*DashboardTemplatingBuilder)) {
	// TODO: support other types of variables: custom, text box, constant, interval, adhoc filters
	fn(&DashboardTemplatingBuilder{dashboard: b.dashboard})
}

func (b *DashboardBuilder) Panels(fn func(*DashboardPanelsBuilder)) {
	// TODO: support other types of panels: time series, singlestat, gauge, bar gauge, table
	// TODO: text, heatmap, alert list, dashboard list, news, pie chart, worldmap panel, logs,
	// TODO: node graph, plugin list
	fn(&DashboardPanelsBuilder{parent: b, panels: &b.dashboard.Panels})
}

func (b *DashboardBuilder) JSON(customJSON interface{}) {
	b.dashboard.CustomJSON.Add(customJSON)
}

func (b *DashboardLinksBuilder) AddExternalLink(fn func(*DashboardLinkExternalBuilder)) {
	linkBuilder := newDashboardLinkExternalBuilder()
	fn(linkBuilder)
	b.dashboard.Links = append(b.dashboard.Links, linkBuilder.link)
}

func (b *DashboardLinksBuilder) AddDashboardsLink(fn func(*DashboardLinkDashboardsBuilder)) {
	linkBuilder := newDashboardLinkDashboardsBuilder()
	fn(linkBuilder)
	b.dashboard.Links = append(b.dashboard.Links, linkBuilder.link)
}

func (b *DashboardAnnotationsBuilder) AddAnnotation(fn func(*DashboardAnnotationBuilder)) {
	annotationBuilder := newDashboardAnnotationBuilder()
	fn(annotationBuilder)
	b.dashboard.Annotations = append(b.dashboard.Annotations, annotationBuilder.annotation)
}

func (b *DashboardTemplatingBuilder) AddQueryVariable(fn func(*DashboardVariableQueryBuilder)) {
	variableBuilder := newDashboardVariableQueryBuilder()
	fn(variableBuilder)
	b.dashboard.Templating = append(b.dashboard.Templating, variableBuilder.variable)
}

func (b *DashboardTemplatingBuilder) AddDatasourceVariable(fn func(*DashboardVariableDatasourceBuilder)) {
	variableBuilder := newDashboardVariableDatasourceBuilder()
	fn(variableBuilder)
	b.dashboard.Templating = append(b.dashboard.Templating, variableBuilder.variable)
}

func (b *DashboardTemplatingBuilder) AddCustomVariable(fn func(*DashboardVariableCustomBuilder)) {
	variableBuilder := newDashboardVariableCustomBuilder()
	fn(variableBuilder)
	b.dashboard.Templating = append(b.dashboard.Templating, variableBuilder.variable)
}

func (b *DashboardPanelsBuilder) AddTablePanel(fn func(*DashboardPanelTableBuilder)) {
	panelBuilder := newDashboardPanelTableBuilder()
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) AddTimeseriesPanel(fn func(builder *DashboardPanelTimeseriesBuilder)) {
	panelBuilder := newDashboardPanelTimeseriesBuilder()
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) AddGraphPanel(fn func(*DashboardPanelGraphBuilder)) {
	panelBuilder := newDashboardPanelGraphBuilder()
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) AddTextPanel(fn func(*DashboardPanelTextBuilder)) {
	panelBuilder := newDashboardPanelTextBuilder()
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) AddGaugePanel(fn func(*DashboardPanelGaugeBuilder)) {
	panelBuilder := newDashboardPanelGaugeBuilder()
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) AddRowPanel(fn func(*DashboardPanelRowBuilder)) {
	panelBuilder := newDashboardPanelRowBuilder(b.parent)
	fn(panelBuilder)

	b.appendPanel(panelBuilder.panel, &panelBuilder.panel.DashboardPanelCommonFields)
}

func (b *DashboardPanelsBuilder) appendPanel(panel schema.DashboardPanel, panelCommonFields *schema.DashboardPanelCommonFields) {
	// Try to keep autogenerated panel IDs unique for persistent panel URLs to be possible.
	if panelCommonFields.ID == 0 {
		hashBytes := sha256.Sum256([]byte(panelCommonFields.Title))
		panelID := int((uint32(hashBytes[0]) | (uint32(hashBytes[1]) << 8) | (uint32(hashBytes[2]) << 16) | (uint32(hashBytes[3]) << 24)) & 0x7FFFFFFF)
		for b.parent.usedPanelIDs[panelID] {
			panelID += 1
		}

		b.parent.usedPanelIDs[panelID] = true
		panelCommonFields.ID = panelID
	} else {
		if b.parent.usedPanelIDs[panelCommonFields.ID] {
			panic("duplicate panel ID detected")
		}

		b.parent.usedPanelIDs[panelCommonFields.ID] = true
	}

	*b.panels = append(*b.panels, panel)
}
