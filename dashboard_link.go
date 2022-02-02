package ggg

import "github.com/g7r/ggg/schema"

type dashboardLinkBuilder struct {
	link *schema.DashboardLinkCommonFields
}

func (b dashboardLinkBuilder) Title(title string) {
	b.link.Title = title
}

func (b dashboardLinkBuilder) IncludeVars(includeVars bool) {
	b.link.IncludeVars = includeVars
}

func (b dashboardLinkBuilder) KeepTime(keepTime bool) {
	b.link.KeepTime = keepTime
}

func (b dashboardLinkBuilder) TargetBlank(targetBlank bool) {
	b.link.TargetBlank = targetBlank
}

func (b dashboardLinkBuilder) OpenInNewTab(newTab bool) {
	b.TargetBlank(newTab)
}
