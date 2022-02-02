package schema

type DashboardAnnotation struct {
	CustomJSON customJSON `json:"-"`

	// Datasource to use for annotation.
	Datasource string `json:"datasource"`
	// Whether annotation is enabled.
	Enable bool `json:"enable"`
	// Whether to hide annotation.
	Hide bool `json:"hide,omitempty"`
	// Annotation icon color.
	IconColor string `json:"iconColor,omitempty"`
	// Name of annotation.
	Name string `json:"name,omitempty"`
	// UNKNOWN
	Type string `json:"type"`
	// Query for annotation data.
	RawQuery string `json:"rawQuery,omitempty"`
	// UNKNOWN
	ShowIn int `json:"showIn"`
	// UNKNOWN
	Limit int `json:"limit"`
}

func (d *DashboardAnnotation) MarshalJSON() ([]byte, error) {
	type plainAnnotation DashboardAnnotation
	return marshalResource((*plainAnnotation)(d), d.CustomJSON)
}
