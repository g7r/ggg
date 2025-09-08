package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelGaugeFieldConfigBuilder struct {
	fieldConfig *schema.DashboardPanelFieldConfig[schema.DashboardPanelGaugeCustom]
}

type DashboardPanelGaugeFieldConfigDefaultsBuilder struct {
	dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelGaugeCustom]
	custom *schema.DashboardPanelGaugeCustom
}

func (b *DashboardPanelGaugeFieldConfigBuilder) Defaults(fn func(builder *DashboardPanelGaugeFieldConfigDefaultsBuilder)) {
	fn(&DashboardPanelGaugeFieldConfigDefaultsBuilder{
		dashboardPanelFieldConfigDefaultsBuilder: dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelGaugeCustom]{
			defaults: &b.fieldConfig.Defaults,
		},
		custom: b.fieldConfig.Defaults.Custom,
	})
}

func (b *DashboardPanelGaugeFieldConfigBuilder) AddOverride(fn func(builder *DashboardPanelFieldConfigOverrideBuilder)) {
	var override schema.DashboardPanelFieldConfigOverride
	fn(&DashboardPanelFieldConfigOverrideBuilder{override: &override})
	b.fieldConfig.Overrides = append(b.fieldConfig.Overrides, override)
}

func (b *DashboardPanelGaugeFieldConfigBuilder) JSON(customJSON interface{}) {
	b.fieldConfig.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGaugeFieldConfigDefaultsBuilder) DefaultsJSON(customJSON interface{}) {
	b.defaults.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGaugeFieldConfigDefaultsBuilder) CustomJSON(customJSON interface{}) {
	b.custom.CustomJSON.Add(customJSON)
}
