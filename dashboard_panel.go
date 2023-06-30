package ggg

import "github.com/g7r/ggg/schema"

type dashboardPanelBuilder struct {
	panel *schema.DashboardPanelCommonFields
}

func (b *dashboardPanelBuilder) ID(id int) {
	b.panel.ID = id
}

func (b *dashboardPanelBuilder) Title(title string) {
	b.panel.Title = title
}

func (b *dashboardPanelBuilder) GridPos(x, y, width, height int) {
	b.panel.GridPos.X = x
	b.panel.GridPos.Y = y
	b.panel.GridPos.Width = width
	b.panel.GridPos.Height = height
}

func (b *dashboardPanelBuilder) Transparent(transparent bool) {
	b.panel.Transparent = transparent
}
