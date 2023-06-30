package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTargetBuilder struct {
	target *schema.DashboardPanelTarget
}

func (b *DashboardPanelTargetBuilder) RefID(refID string) {
	b.target.RefID = refID
}

func (b *DashboardPanelTargetBuilder) Format(format schema.TargetFormat) {
	b.target.Format = format
}

func (b *DashboardPanelTargetBuilder) Hide(hide bool) {
	b.target.Hide = hide
}

func (b *DashboardPanelTargetBuilder) Instant(instant bool) {
	b.target.Instant = instant
}

func (b *DashboardPanelTargetBuilder) Exemplar(exemplar bool) {
	b.target.Exemplar = exemplar
}

func (b *DashboardPanelTargetBuilder) Expr(expr string) {
	b.target.Expr = expr
}

func (b *DashboardPanelTargetBuilder) Interval(interval string) {
	b.target.Interval = interval
}

func (b *DashboardPanelTargetBuilder) IntervalFactor(intervalFactor int) {
	b.target.IntervalFactor = intervalFactor
}

func (b *DashboardPanelTargetBuilder) LegendFormat(legendFormat string) {
	b.target.LegendFormat = legendFormat
}

func (b *DashboardPanelTargetBuilder) JSON(customJSON interface{}) {
	b.target.CustomJSON.Add(customJSON)
}
