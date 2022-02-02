package ggg

import "github.com/g7r/ggg/schema"

type DashboardTimepickerBuilder struct {
	timepicker *schema.DashboardTimepicker
}

func (b *DashboardTimepickerBuilder) Hidden(hidden bool) {
	b.timepicker.Hidden = hidden
}

func (b *DashboardTimepickerBuilder) RefreshIntervals(intervals ...string) {
	b.timepicker.RefreshIntervals = intervals
}

func (b *DashboardTimepickerBuilder) NowDelay(nowDelay string) {
	b.timepicker.NowDelay = nowDelay
}
