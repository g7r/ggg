package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTableFieldConfigBuilder struct {
	fieldConfig *schema.DashboardPanelFieldConfig[schema.DashboardPanelTableCustom]
}

type DashboardPanelTableFieldConfigDefaultsBuilder struct {
	dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelTableCustom]
	custom *schema.DashboardPanelTableCustom
}

func (b *DashboardPanelTableFieldConfigBuilder) Defaults(fn func(builder *DashboardPanelTableFieldConfigDefaultsBuilder)) {
	fn(&DashboardPanelTableFieldConfigDefaultsBuilder{
		dashboardPanelFieldConfigDefaultsBuilder: dashboardPanelFieldConfigDefaultsBuilder[schema.DashboardPanelTableCustom]{
			defaults: &b.fieldConfig.Defaults,
		},
		custom: b.fieldConfig.Defaults.Custom,
	})
}

func (b *DashboardPanelTableFieldConfigBuilder) AddOverride(fn func(builder *DashboardPanelFieldConfigOverrideBuilder)) {
	var override schema.DashboardPanelFieldConfigOverride
	fn(&DashboardPanelFieldConfigOverrideBuilder{override: &override})
	b.fieldConfig.Overrides = append(b.fieldConfig.Overrides, override)
}

func (b *DashboardPanelTableFieldConfigBuilder) JSON(customJSON interface{}) {
	b.fieldConfig.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTableFieldConfigDefaultsBuilder) DefaultsJSON(customJSON interface{}) {
	b.defaults.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelTableFieldConfigDefaultsBuilder) getCustom() *schema.DashboardPanelTableCustom {
	if b.custom == nil {
		b.custom = &schema.DashboardPanelTableCustom{}
		b.defaults.Custom = b.custom
	}

	return b.custom
}

func (b *DashboardPanelTableFieldConfigDefaultsBuilder) CustomJSON(customJSON interface{}) {
	b.getCustom().CustomJSON.Add(customJSON)
}
