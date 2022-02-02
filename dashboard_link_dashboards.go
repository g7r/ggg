package ggg

import "github.com/g7r/ggg/schema"

type DashboardLinkDashboardsBuilder struct {
	dashboardLinkBuilder
	link *schema.DashboardLinkDashboards
}

func newDashboardLinkDashboardsBuilder() *DashboardLinkDashboardsBuilder {
	link := schema.DashboardLinkDashboards{
		DashboardLinkCommonFields: schema.DashboardLinkCommonFields{
			Type: "dashboards",
		},
	}

	return &DashboardLinkDashboardsBuilder{
		dashboardLinkBuilder: dashboardLinkBuilder{link: &link.DashboardLinkCommonFields},
		link:                 &link,
	}
}

func (b *DashboardLinkDashboardsBuilder) AsDropdown(asDropdown bool) {
	b.link.AsDropdown = asDropdown
}

func (b *DashboardLinkDashboardsBuilder) Tags(tags ...string) {
	b.link.Tags = tags
}

func (b *DashboardLinkDashboardsBuilder) JSON(customJSON interface{}) {
	b.link.CustomJSON.Add(customJSON)
}
