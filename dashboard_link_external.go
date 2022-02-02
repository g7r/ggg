package ggg

import "github.com/g7r/ggg/schema"

type DashboardLinkExternalBuilder struct {
	dashboardLinkBuilder
	link *schema.DashboardLinkExternal
}

func newDashboardLinkExternalBuilder() *DashboardLinkExternalBuilder {
	link := schema.DashboardLinkExternal{
		DashboardLinkCommonFields: schema.DashboardLinkCommonFields{
			Type: "link",
		},
		Icon: schema.LinkIconExternalLink,
	}

	return &DashboardLinkExternalBuilder{
		dashboardLinkBuilder: dashboardLinkBuilder{link: &link.DashboardLinkCommonFields},
		link:                 &link,
	}
}

func (b *DashboardLinkExternalBuilder) URL(url string) {
	b.link.URL = url
}

func (b *DashboardLinkExternalBuilder) Icon(icon schema.LinkIcon) {
	b.link.Icon = icon
}

func (b *DashboardLinkExternalBuilder) Tooltip(tooltip string) {
	b.link.Tooltip = tooltip
}

func (b *DashboardLinkExternalBuilder) JSON(customJSON interface{}) {
	b.link.CustomJSON.Add(customJSON)
}
