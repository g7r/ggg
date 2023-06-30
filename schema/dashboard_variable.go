package schema

type DashboardVariableCommonFields struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Label string `json:"label"`
	// Do not forward variable in links
	SkipURLSync bool `json:"skipUrlSync"`
	// Human readable definition, usually mirrors query
	Definition string `json:"definition,omitempty"`
	// Tooltip text
	Description string `json:"description,omitempty"`
	// Hide label or entire variable from user
	Hide DashboardVariableHide `json:"hide"`

	Current *DashboardVariableCurrent `json:"current,omitempty"`
}

type DashboardVariableQuery struct {
	CustomJSON customJSON `json:"-"`

	DashboardVariable             `json:"-"`
	DashboardVariableCommonFields `json:",inline"`

	Datasource string `json:"datasource,omitempty"`
	// Enable 'All' option in result value list
	IncludeAll bool `json:"includeAll"`
	// Use this value as a custom 'All' value
	AllValue string `json:"allValue,omitempty"`
	// Allow multiple values to be selected
	Multi bool `json:"multi"`
	// When to refresh variable values
	Refresh DashboardVariableRefresh `json:"refresh"`
	// Regular expression to filter query results
	Regex string `json:"regex"`
	// Result sort order
	Sort DashboardVariableSort `json:"sort"`

	Query string `json:"query"`
}

type DashboardVariableDatasource struct {
	CustomJSON customJSON `json:"-"`

	DashboardVariable             `json:"-"`
	DashboardVariableCommonFields `json:",inline"`

	// When to refresh variable values
	Refresh DashboardVariableRefresh `json:"refresh"`
	// Enable 'All' option in result value list
	IncludeAll bool `json:"includeAll"`
	// Use this value as a custom 'All' value
	AllValue string `json:"allValue,omitempty"`
	// Allow multiple values to be selected
	Multi bool `json:"multi"`

	Query string `json:"query"`
}

type DashboardVariableCustom struct {
	CustomJSON customJSON `json:"-"`

	DashboardVariable             `json:"-"`
	DashboardVariableCommonFields `json:",inline"`

	// When to refresh variable values
	Refresh DashboardVariableRefresh `json:"refresh"`
	// Enable 'All' option in result value list
	IncludeAll bool `json:"includeAll"`
	// Use this value as a custom 'All' value
	AllValue string `json:"allValue,omitempty"`
	// Allow multiple values to be selected
	Multi bool `json:"multi"`

	Query string `json:"query"`
}

type DashboardVariableCurrent struct {
	CustomJSON customJSON `json:"-"`

	// True if the value was selected by user
	Selected bool `json:"selected"`
	// Text to show
	Text string `json:"text"`
	// The real value. There are possible special values like "$__auto_interval" or "$__all".
	Value string `json:"value"`
}

func (d *DashboardVariableQuery) MarshalJSON() ([]byte, error) {
	type plainVariable DashboardVariableQuery
	return marshalResource((*plainVariable)(d), d.CustomJSON)
}

func (d *DashboardVariableDatasource) MarshalJSON() ([]byte, error) {
	type plainVariable DashboardVariableDatasource
	return marshalResource((*plainVariable)(d), d.CustomJSON)
}

func (d *DashboardVariableCurrent) MarshalJSON() ([]byte, error) {
	type plainCurrent DashboardVariableCurrent
	return marshalResource((*plainCurrent)(d), d.CustomJSON)
}

func (d *DashboardVariableCustom) MarshalJSON() ([]byte, error) {
	type plainCustom DashboardVariableCustom
	return marshalResource((*plainCustom)(d), d.CustomJSON)
}
