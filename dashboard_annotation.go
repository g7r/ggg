package ggg

import (
	"fmt"

	"github.com/g7r/ggg/schema"
)

type DashboardAnnotationBuilder struct {
	annotation *schema.DashboardAnnotation
}

func newDashboardAnnotationBuilder() *DashboardAnnotationBuilder {
	return &DashboardAnnotationBuilder{annotation: &schema.DashboardAnnotation{
		Type:   "dashboard",
		Enable: true,
	}}
}

func (b *DashboardAnnotationBuilder) Name(name string) {
	b.annotation.Name = name
}

func (b *DashboardAnnotationBuilder) Datasource(ds string) {
	b.annotation.Datasource = ds
}

func (b *DashboardAnnotationBuilder) DatasourceVar(dashboardVar string) {
	b.annotation.Datasource = fmt.Sprintf("${%s}", dashboardVar)
}

func (b *DashboardAnnotationBuilder) Enable(enable bool) {
	b.annotation.Enable = enable
}

func (b *DashboardAnnotationBuilder) Hide(hide bool) {
	b.annotation.Hide = hide
}

func (b *DashboardAnnotationBuilder) IconColor(iconColor string) {
	b.annotation.IconColor = iconColor
}

func (b *DashboardAnnotationBuilder) Type(type_ string) {
	b.annotation.Type = type_
}

func (b *DashboardAnnotationBuilder) RawQuery(rawQuery string) {
	b.annotation.RawQuery = rawQuery
}

func (b *DashboardAnnotationBuilder) Expr(expr string) {
	b.annotation.Expr = expr
}

func (b *DashboardAnnotationBuilder) Step(step string) {
	b.annotation.Step = step
}

func (b *DashboardAnnotationBuilder) ShowIn(showIn int) {
	b.annotation.ShowIn = showIn
}

func (b *DashboardAnnotationBuilder) Limit(limit int) {
	b.annotation.Limit = limit
}

func (b *DashboardAnnotationBuilder) JSON(customJSON interface{}) {
	b.annotation.CustomJSON.Add(customJSON)
}
