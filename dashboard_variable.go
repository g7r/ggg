package ggg

import (
	"github.com/g7r/ggg/schema"
)

type dashboardVariableBuilder struct {
	variable *schema.DashboardVariableCommonFields
}

func (b dashboardVariableBuilder) Name(name string) {
	b.variable.Name = name
}

func (b dashboardVariableBuilder) Label(label string) {
	b.variable.Label = label
}

func (b dashboardVariableBuilder) SkipURLSync(skip bool) {
	b.variable.SkipURLSync = skip
}

func (b dashboardVariableBuilder) Definition(definition string) {
	b.variable.Definition = definition
}

func (b dashboardVariableBuilder) Description(description string) {
	b.variable.Description = description
}

func (b dashboardVariableBuilder) Hide(hide schema.DashboardVariableHide) {
	b.variable.Hide = hide
}

func (b dashboardVariableBuilder) Current(fn func(*DashboardVariableCurrentBuilder)) {
	currentBuilder := newDashboardVariableCurrentBuilder()
	fn(currentBuilder)
	b.variable.Current = currentBuilder.current
}
