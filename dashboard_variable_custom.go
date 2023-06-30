package ggg

import "github.com/g7r/ggg/schema"

type DashboardVariableCustomBuilder struct {
	dashboardVariableBuilder
	variable *schema.DashboardVariableCustom
}

func newDashboardVariableCustomBuilder() *DashboardVariableCustomBuilder {
	variable := schema.DashboardVariableCustom{
		DashboardVariableCommonFields: schema.DashboardVariableCommonFields{
			Type: "custom",
		},
	}

	return &DashboardVariableCustomBuilder{
		dashboardVariableBuilder: dashboardVariableBuilder{variable: &variable.DashboardVariableCommonFields},
		variable:                 &variable,
	}
}

func (b *DashboardVariableCustomBuilder) Multi(multi bool) {
	b.variable.Multi = multi
}

func (b *DashboardVariableCustomBuilder) Query(query string) {
	b.variable.Query = query
}

func (b DashboardVariableCustomBuilder) JSON(customJSON interface{}) {
	b.variable.CustomJSON.Add(customJSON)
}
