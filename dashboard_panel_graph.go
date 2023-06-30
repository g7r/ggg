package ggg

import (
	"fmt"

	"github.com/g7r/ggg/schema"
)

type DashboardPanelGraphBuilder struct {
	dashboardPanelBuilder
	panel *schema.DashboardPanelGraph

	xAxisDefined bool
	nextRefID    rune
}

type DashboardPanelGraphXAxisBuilder struct {
	parent *DashboardPanelGraphBuilder
}

type dashboardPanelGraphXAxisBuilder struct {
	axis *schema.DashboardPanelGraphXAxisCommonFields
}

type DashboardPanelGraphXAxisTimeBuilder struct {
	dashboardPanelGraphXAxisBuilder
	axis *schema.DashboardPanelGraphXAxisTime
}

type DashboardPanelGraphXAxisSeriesBuilder struct {
	dashboardPanelGraphXAxisBuilder
	axis *schema.DashboardPanelGraphXAxisSeries
}

type DashboardPanelGraphXAxisHistogramBuilder struct {
	dashboardPanelGraphXAxisBuilder
	axis *schema.DashboardPanelGraphXAxisHistogram
}

type DashboardPanelGraphTooltipBuilder struct {
	tooltip *schema.DashboardPanelGraphTooltip
}

type DashboardPanelGraphYAxisOptionsBuilder struct {
	options *schema.DashboardPanelGraphYAxisOptions
}

type DashboardPanelGraphLegendBuilder struct {
	legend *schema.DashboardPanelGraphLegend
}

type DashboardPanelGraphYAxisBuilder struct {
	yaxis *schema.DashboardPanelGraphYAxis
}

type DashboardPanelGraphTargetsBuilder struct {
	parent *DashboardPanelGraphBuilder
}

func newDashboardPanelGraphBuilder() *DashboardPanelGraphBuilder {
	panel := schema.DashboardPanelGraph{
		DashboardPanelCommonFields: schema.DashboardPanelCommonFields{
			Type: "graph",
		},
		PluginVersion: "7.4.5",
		Renderer:      "flot",
		NullPointMode: schema.NullPointModeNull,
		Lines:         true,
		LineWidth:     1,
		Fill:          1,
		PointRadius:   5,
		DashLength:    10,
		SpaceLength:   10,
		XAxis: &schema.DashboardPanelGraphXAxisTime{
			DashboardPanelGraphXAxisCommonFields: schema.DashboardPanelGraphXAxisCommonFields{
				Show: true,
				Mode: "time",
			},
		},
		Tooltip: schema.DashboardPanelGraphTooltip{
			ValueType: "individual",
			Shared:    true,
		},
		Options: schema.DashboardPanelGraphOptions{
			AlertThreshold: true,
		},
		Legend: schema.DashboardPanelGraphLegend{
			Show: true,
		},
		YAxes: schema.DashboardPanelGraphYAxes{
			Left: schema.DashboardPanelGraphYAxis{
				Show:    true,
				Format:  schema.UnitMiscShort,
				LogBase: 1.0,
			},
			Right: schema.DashboardPanelGraphYAxis{
				Show:    true,
				Format:  schema.UnitMiscShort,
				LogBase: 1.0,
			},
		},
	}

	return &DashboardPanelGraphBuilder{
		dashboardPanelBuilder: dashboardPanelBuilder{panel: &panel.DashboardPanelCommonFields},
		panel:                 &panel,
		nextRefID:             'A',
	}
}

func newDashboardPanelGraphXAxisTimeBuilder(axis *schema.DashboardPanelGraphXAxisTime) *DashboardPanelGraphXAxisTimeBuilder {
	return &DashboardPanelGraphXAxisTimeBuilder{
		dashboardPanelGraphXAxisBuilder: dashboardPanelGraphXAxisBuilder{
			axis: &axis.DashboardPanelGraphXAxisCommonFields,
		},
		axis: axis,
	}
}

func newDashboardPanelGraphXAxisSeriesBuilder(axis *schema.DashboardPanelGraphXAxisSeries) *DashboardPanelGraphXAxisSeriesBuilder {
	return &DashboardPanelGraphXAxisSeriesBuilder{
		dashboardPanelGraphXAxisBuilder: dashboardPanelGraphXAxisBuilder{
			axis: &axis.DashboardPanelGraphXAxisCommonFields,
		},
		axis: axis,
	}
}

func newDashboardPanelGraphXAxisHistogramBuilder(axis *schema.DashboardPanelGraphXAxisHistogram) *DashboardPanelGraphXAxisHistogramBuilder {
	return &DashboardPanelGraphXAxisHistogramBuilder{
		dashboardPanelGraphXAxisBuilder: dashboardPanelGraphXAxisBuilder{
			axis: &axis.DashboardPanelGraphXAxisCommonFields,
		},
		axis: axis,
	}
}

func (b *DashboardPanelGraphBuilder) Renderer(renderer string) {
	b.panel.Renderer = renderer
}

func (b *DashboardPanelGraphBuilder) DatasourceGlobal(datasource string) {
	b.panel.Datasource = datasource
}

func (b *DashboardPanelGraphBuilder) DatasourceVar(datasourceVar string) {
	b.panel.Datasource = fmt.Sprintf("${%s}", datasourceVar)
}

func (b *DashboardPanelGraphBuilder) Bars(bars bool) {
	b.panel.Bars = bars
}

func (b *DashboardPanelGraphBuilder) Lines(lines bool) {
	b.panel.Lines = lines
}

func (b *DashboardPanelGraphBuilder) Dashes(dashes bool) {
	b.panel.Dashes = dashes
}

func (b *DashboardPanelGraphBuilder) Points(points bool) {
	b.panel.Points = points
}

func (b *DashboardPanelGraphBuilder) Percentage(percentage bool) {
	b.panel.Percentage = percentage
}

func (b *DashboardPanelGraphBuilder) NullPointMode(mode schema.NullPointMode) {
	b.panel.NullPointMode = mode
}

func (b *DashboardPanelGraphBuilder) Decimals(decimals int) {
	b.panel.Decimals = &decimals
}

func (b *DashboardPanelGraphBuilder) DashLength(dashLength int) {
	b.panel.DashLength = dashLength
}

func (b *DashboardPanelGraphBuilder) Fill(fill int) {
	b.panel.Fill = fill
}

func (b *DashboardPanelGraphBuilder) FillGradient(fillGradient int) {
	b.panel.FillGradient = fillGradient
}

func (b *DashboardPanelGraphBuilder) LineWidth(lineWidth int) {
	b.panel.LineWidth = lineWidth
}

func (b *DashboardPanelGraphBuilder) PointRadius(pointRadius float64) {
	b.panel.PointRadius = pointRadius
}

func (b *DashboardPanelGraphBuilder) HiddenSeries(hiddenSeries bool) {
	b.panel.HiddenSeries = hiddenSeries
}

func (b *DashboardPanelGraphBuilder) Stack(stack bool) {
	b.panel.Stack = stack
}

func (b *DashboardPanelGraphBuilder) AlertThreshold(alertThreshold bool) {
	b.panel.Options.AlertThreshold = alertThreshold
}

func (b *DashboardPanelGraphBuilder) XAxis(fn func(*DashboardPanelGraphXAxisBuilder)) {
	fn(&DashboardPanelGraphXAxisBuilder{parent: b})
}

func (b *DashboardPanelGraphBuilder) Tooltip(fn func(*DashboardPanelGraphTooltipBuilder)) {
	fn(&DashboardPanelGraphTooltipBuilder{tooltip: &b.panel.Tooltip})
}

func (b *DashboardPanelGraphBuilder) YAxisLeft(fn func(builder *DashboardPanelGraphYAxisBuilder)) {
	fn(&DashboardPanelGraphYAxisBuilder{yaxis: &b.panel.YAxes.Left})
}

func (b *DashboardPanelGraphBuilder) YAxisRight(fn func(builder *DashboardPanelGraphYAxisBuilder)) {
	fn(&DashboardPanelGraphYAxisBuilder{yaxis: &b.panel.YAxes.Right})
}

func (b *DashboardPanelGraphBuilder) Legend(fn func(builder *DashboardPanelGraphLegendBuilder)) {
	fn(&DashboardPanelGraphLegendBuilder{legend: &b.panel.Legend})
}

func (b *DashboardPanelGraphBuilder) Targets(fn func(builder *DashboardPanelGraphTargetsBuilder)) {
	fn(&DashboardPanelGraphTargetsBuilder{parent: b})
}

func (b *DashboardPanelGraphBuilder) JSON(customJSON interface{}) {
	b.panel.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphXAxisBuilder) Time(fn func(*DashboardPanelGraphXAxisTimeBuilder)) {
	var axis *schema.DashboardPanelGraphXAxisTime

	if !b.parent.xAxisDefined {
		var ok bool
		axis, ok = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisTime)
		if !ok {
			axis = &schema.DashboardPanelGraphXAxisTime{
				DashboardPanelGraphXAxisCommonFields: schema.DashboardPanelGraphXAxisCommonFields{
					Show: true,
					Mode: "time",
				},
			}
		}
	} else {
		axis = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisTime)
	}

	fn(newDashboardPanelGraphXAxisTimeBuilder(axis))
	b.parent.panel.XAxis = axis
	b.parent.xAxisDefined = true
}

func (b *DashboardPanelGraphXAxisBuilder) Series(fn func(*DashboardPanelGraphXAxisSeriesBuilder)) {
	var axis *schema.DashboardPanelGraphXAxisSeries

	if !b.parent.xAxisDefined {
		var ok bool
		axis, ok = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisSeries)
		if !ok {
			axis = &schema.DashboardPanelGraphXAxisSeries{
				DashboardPanelGraphXAxisCommonFields: schema.DashboardPanelGraphXAxisCommonFields{
					Show: true,
					Mode: "series",
				},
				Values: []schema.XAxisValue{schema.XAxisValueTotal},
			}
		}
	} else {
		axis = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisSeries)
	}

	fn(newDashboardPanelGraphXAxisSeriesBuilder(axis))
	b.parent.panel.XAxis = axis
	b.parent.xAxisDefined = true
}

func (b *DashboardPanelGraphXAxisBuilder) Histogram(fn func(*DashboardPanelGraphXAxisHistogramBuilder)) {
	var axis *schema.DashboardPanelGraphXAxisHistogram

	if !b.parent.xAxisDefined {
		var ok bool
		axis, ok = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisHistogram)
		if !ok {
			axis = &schema.DashboardPanelGraphXAxisHistogram{
				DashboardPanelGraphXAxisCommonFields: schema.DashboardPanelGraphXAxisCommonFields{
					Show: true,
					Mode: "histogram",
				},
			}
		}
	} else {
		axis = b.parent.panel.XAxis.(*schema.DashboardPanelGraphXAxisHistogram)
	}

	fn(newDashboardPanelGraphXAxisHistogramBuilder(axis))
	b.parent.panel.XAxis = axis
	b.parent.xAxisDefined = true
}

func (b *dashboardPanelGraphXAxisBuilder) Show(show bool) {
	b.axis.Show = show
}

func (b *dashboardPanelGraphXAxisBuilder) Mode(mode string) {
	b.axis.Mode = mode
}

func (b *DashboardPanelGraphXAxisTimeBuilder) JSON(customJSON interface{}) {
	b.axis.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphXAxisSeriesBuilder) Value(value schema.XAxisValue) {
	b.axis.Values = []schema.XAxisValue{value}
}

func (b *DashboardPanelGraphXAxisSeriesBuilder) JSON(customJSON interface{}) {
	b.axis.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphXAxisHistogramBuilder) Buckets(buckets int) {
	b.axis.Buckets = buckets
}

func (b *DashboardPanelGraphXAxisHistogramBuilder) Min(min float64) {
	b.axis.Min = min
}

func (b *DashboardPanelGraphXAxisHistogramBuilder) Max(max float64) {
	b.axis.Max = max
}

func (b *DashboardPanelGraphXAxisHistogramBuilder) JSON(customJSON interface{}) {
	b.axis.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphTooltipBuilder) Sort(sort schema.GraphTooltipSort) {
	b.tooltip.Sort = sort
}

func (b *DashboardPanelGraphTooltipBuilder) Shared(shared bool) {
	b.tooltip.Shared = shared
}

func (b *DashboardPanelGraphTooltipBuilder) JSON(customJSON interface{}) {
	b.tooltip.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphYAxisOptionsBuilder) Align(align bool) {
	b.options.Align = align
}

func (b *DashboardPanelGraphYAxisOptionsBuilder) AlignLevel(alignLevel float64) {
	b.options.AlignLevel = alignLevel
}

func (b *DashboardPanelGraphYAxisOptionsBuilder) JSON(customJSON interface{}) {
	b.options.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphLegendBuilder) Show(show bool) {
	b.legend.Show = show
}

func (b *DashboardPanelGraphLegendBuilder) Values(values bool) {
	b.legend.Values = values
}

func (b *DashboardPanelGraphLegendBuilder) Avg(avg bool) {
	b.legend.Avg = avg
	b.legend.Values = b.legend.Values || avg
}

func (b *DashboardPanelGraphLegendBuilder) Current(current bool) {
	b.legend.Current = current
	b.legend.Values = b.legend.Values || current
}

func (b *DashboardPanelGraphLegendBuilder) Min(min bool) {
	b.legend.Min = min
	b.legend.Values = b.legend.Values || min
}

func (b *DashboardPanelGraphLegendBuilder) Max(max bool) {
	b.legend.Max = max
	b.legend.Values = b.legend.Values || max
}

func (b *DashboardPanelGraphLegendBuilder) Total(total bool) {
	b.legend.Total = total
	b.legend.Values = b.legend.Values || total
}

func (b *DashboardPanelGraphLegendBuilder) AlignAsTable(alignAsTable bool) {
	b.legend.AlignAsTable = alignAsTable
}

func (b *DashboardPanelGraphLegendBuilder) RightSide(rightSide bool) {
	b.legend.RightSide = rightSide
}

func (b *DashboardPanelGraphLegendBuilder) SideWidth(sideWidth float64) {
	b.legend.SideWidth = sideWidth
}

func (b *DashboardPanelGraphLegendBuilder) Sort(sort string) {
	b.legend.Sort = sort
}

func (b *DashboardPanelGraphLegendBuilder) SortDesc(sortDesc bool) {
	b.legend.SortDesc = sortDesc
}

func (b *DashboardPanelGraphLegendBuilder) JSON(customJSON interface{}) {
	b.legend.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphYAxisBuilder) Format(format schema.Unit) {
	b.yaxis.Format = format
}

func (b *DashboardPanelGraphYAxisBuilder) Show(show bool) {
	b.yaxis.Show = show
}

func (b *DashboardPanelGraphYAxisBuilder) Label(label string) {
	b.yaxis.Label = label
}

func (b *DashboardPanelGraphYAxisBuilder) LogBase(logBase float64) {
	b.yaxis.LogBase = logBase
}

func (b *DashboardPanelGraphYAxisBuilder) Min(min float64) {
	b.yaxis.Min = &min
}

func (b *DashboardPanelGraphYAxisBuilder) Max(max float64) {
	b.yaxis.Max = &max
}

func (b *DashboardPanelGraphYAxisBuilder) Decimals(decimals int) {
	b.yaxis.Decimals = &decimals
}

func (b *DashboardPanelGraphYAxisBuilder) JSON(customJSON interface{}) {
	b.yaxis.CustomJSON.Add(customJSON)
}

func (b *DashboardPanelGraphTargetsBuilder) AddTarget(fn func(builder *DashboardPanelTargetBuilder)) {
	refID := string([]rune{b.parent.nextRefID})
	b.parent.nextRefID += 1

	target := schema.DashboardPanelTarget{
		RefID:          refID,
		Format:         schema.TargetFormatTimeSeries,
		IntervalFactor: 1,
	}

	fn(&DashboardPanelTargetBuilder{target: &target})
	b.parent.panel.Targets = append(b.parent.panel.Targets, &target)
}
