package ggg

import "github.com/g7r/ggg/schema"

type DashboardVariableCurrentBuilder struct {
	current *schema.DashboardVariableCurrent
}

func newDashboardVariableCurrentBuilder() *DashboardVariableCurrentBuilder {
	return &DashboardVariableCurrentBuilder{current: &schema.DashboardVariableCurrent{}}
}

func (b *DashboardVariableCurrentBuilder) Selected(selected bool) {
	b.current.Selected = selected
}

func (b *DashboardVariableCurrentBuilder) Text(text string) {
	b.current.Text = text
}

func (b *DashboardVariableCurrentBuilder) Value(value string) {
	b.current.Value = value
}

func (b *DashboardVariableCurrentBuilder) ValueAndText(value string) {
	b.current.Value = value
	b.current.Text = value
}

func (b *DashboardVariableCurrentBuilder) ValueAll() {
	b.current.Text = "All"
	b.current.Value = "$__all"
}

func (b *DashboardVariableCurrentBuilder) JSON(customJSON interface{}) {
	b.current.CustomJSON.Add(customJSON)
}
