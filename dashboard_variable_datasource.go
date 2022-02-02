package ggg

import "github.com/g7r/ggg/schema"

type DashboardVariableDatasourceBuilder struct {
	dashboardVariableBuilder
	variable *schema.DashboardVariableDatasource
}

func newDashboardVariableDatasourceBuilder() *DashboardVariableDatasourceBuilder {
	variable := schema.DashboardVariableDatasource{
		DashboardVariableCommonFields: schema.DashboardVariableCommonFields{
			Type: "datasource",
		},
	}

	return &DashboardVariableDatasourceBuilder{
		dashboardVariableBuilder: dashboardVariableBuilder{variable: &variable.DashboardVariableCommonFields},
		variable:                 &variable,
	}
}

func (b *DashboardVariableDatasourceBuilder) IncludeAll(includeAll bool) {
	b.variable.IncludeAll = includeAll
}

func (b *DashboardVariableDatasourceBuilder) AllValue(allValue string) {
	b.variable.AllValue = allValue
}

func (b *DashboardVariableDatasourceBuilder) Multi(multi bool) {
	b.variable.Multi = multi
}

func (b *DashboardVariableDatasourceBuilder) Query(query string) {
	b.variable.Query = query
}

func (b *DashboardVariableDatasourceBuilder) Refresh(refresh schema.DashboardVariableRefresh) {
	b.variable.Refresh = refresh
}

func (b *DashboardVariableDatasourceBuilder) JSON(customJSON interface{}) {
	b.variable.CustomJSON.Add(customJSON)
}
