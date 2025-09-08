package ggg

import "github.com/g7r/ggg/schema"

type DashboardPanelTargetBuilder struct {
	target *schema.DashboardPanelTarget
}

type dashboardPanelTargetQueryBuilder struct {
	customJSON interface{ Add(any) }
	target     *schema.DashboardPanelTargetQuery
}

type dashboardPanelTargetSQLBuilder struct {
	dashboardPanelTargetQueryBuilder
	target *schema.DashboardPanelTargetSQL
}

type DashboardPanelTargetAthenaSQLBuilder struct {
	dashboardPanelTargetQueryBuilder
	target *schema.DashboardPanelTargetAthenaSQL
}

func (b *DashboardPanelTargetBuilder) RefID(refID string) {
	b.target.RefID = refID
}

func (b *DashboardPanelTargetBuilder) Format(format schema.QueryFormat) {
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

func (b *dashboardPanelTargetQueryBuilder) RefID(refID string) {
	b.target.RefID = refID
}

func (b *dashboardPanelTargetQueryBuilder) Hide(hide bool) {
	b.target.Hide = hide
}

func (b *dashboardPanelTargetSQLBuilder) Format(format schema.QueryFormat) {
	b.target.Format = format
}

func (b *dashboardPanelTargetSQLBuilder) RawSQL(rawSQL string) {
	b.target.RawSQL = rawSQL
}

func (b *DashboardPanelTargetAthenaSQLBuilder) Format(format schema.QueryFormatAthenaSQL) {
	b.target.Format = format
}

func (b *DashboardPanelTargetAthenaSQLBuilder) RawSQL(rawSQL string) {
	b.target.RawSQL = rawSQL
}

func (b *dashboardPanelTargetQueryBuilder) JSON(customJSON interface{}) {
	b.customJSON.Add(customJSON)
}

func (b *dashboardPanelTargetSQLBuilder) JSON(customJSON interface{}) {
	b.customJSON.Add(customJSON)
}

func (b *DashboardPanelTargetAthenaSQLBuilder) JSON(customJSON interface{}) {
	b.customJSON.Add(customJSON)
}

func (b *DashboardPanelTargetBuilder) JSON(customJSON interface{}) {
	b.target.CustomJSON.Add(customJSON)
}
