package ggg

import "github.com/g7r/ggg/schema"

type DashboardDatasourceRefBuilder struct {
	parent *schema.DashboardDatasourceRef
}

func (b *DashboardDatasourceRefBuilder) Type(typ string) {
	b.parent.Type = typ
}

func (b *DashboardDatasourceRefBuilder) UID(uid string) {
	b.parent.UID = uid
}

func (b *DashboardDatasourceRefBuilder) Athena(uid string) {
	b.parent.Type = "grafana-athena-datasource"
	b.parent.UID = uid
}
