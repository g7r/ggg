package ggg

import "github.com/g7r/ggg/schema"

type DashboardAnnotationBuilder struct {
	annotation *schema.DashboardAnnotation
}

func newDashboardAnnotationBuilder() *DashboardAnnotationBuilder {
	return &DashboardAnnotationBuilder{annotation: &schema.DashboardAnnotation{
		Type:   "dashboard",
		Enable: true,
	}}
}

func (b *DashboardAnnotationBuilder) Datasource(ds string) {
	b.annotation.Datasource = ds
}

func (b *DashboardAnnotationBuilder) JSON(customJSON interface{}) {
	b.annotation.CustomJSON.Add(customJSON)
}
