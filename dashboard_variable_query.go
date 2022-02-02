package ggg

import (
	"fmt"

	"github.com/g7r/ggg/schema"
)

type DashboardVariableQueryBuilder struct {
	dashboardVariableBuilder
	variable *schema.DashboardVariableQuery
}

func newDashboardVariableQueryBuilder() *DashboardVariableQueryBuilder {
	variable := schema.DashboardVariableQuery{
		DashboardVariableCommonFields: schema.DashboardVariableCommonFields{
			Type: "query",
		},
	}

	return &DashboardVariableQueryBuilder{
		dashboardVariableBuilder: dashboardVariableBuilder{variable: &variable.DashboardVariableCommonFields},
		variable:                 &variable,
	}
}

func (b *DashboardVariableQueryBuilder) DatasourceGlobal(datasource string) {
	b.variable.Datasource = datasource
}

func (b *DashboardVariableQueryBuilder) DatasourceVar(datasourceVar string) {
	b.variable.Datasource = fmt.Sprintf("${%s}", datasourceVar)
}

func (b *DashboardVariableQueryBuilder) IncludeAll(includeAll bool) {
	b.variable.IncludeAll = includeAll
}

func (b *DashboardVariableQueryBuilder) AllValue(allValue string) {
	b.variable.AllValue = allValue
}

func (b *DashboardVariableQueryBuilder) Multi(multi bool) {
	b.variable.Multi = multi
}

func (b *DashboardVariableQueryBuilder) Refresh(refresh schema.DashboardVariableRefresh) {
	b.variable.Refresh = refresh
}

func (b *DashboardVariableQueryBuilder) Regex(regex string) {
	b.variable.Regex = regex
}

func (b *DashboardVariableQueryBuilder) Sort(sort schema.DashboardVariableSort) {
	b.variable.Sort = sort
}

func (b *DashboardVariableQueryBuilder) Query(query string) {
	b.variable.Query = query
}

func (b *DashboardVariableQueryBuilder) JSON(customJSON interface{}) {
	b.variable.CustomJSON.Add(customJSON)
}
