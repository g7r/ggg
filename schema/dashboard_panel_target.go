package schema

type QueryFormat string

const (
	QueryFormatTimeSeries QueryFormat = "time_series"
	QueryFormatTable      QueryFormat = "table"
)

type QueryFormatAthenaSQL int

const (
	QueryFormatAthenaSQLTimeSeries QueryFormatAthenaSQL = 0
	QueryFormatAthenaSQLTable      QueryFormatAthenaSQL = 1
	QueryFormatAthenaSQLLogs       QueryFormatAthenaSQL = 2
)

type DashboardPanelTarget struct {
	CustomJSON customJSON `json:"-"`

	RefID  string      `json:"refId"`
	Format QueryFormat `json:"format"`
	Hide   bool        `json:"hide,omitempty"`
	Expr   string      `json:"expr,omitempty"`
	// Min step (should be like "5s", "15m")
	Interval string `json:"interval,omitempty"`
	// Resolution is 1/IntervalFactor
	IntervalFactor int    `json:"intervalFactor"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Exemplar       bool   `json:"exemplar,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
}

type DashboardPanelGenericTarget interface {
	isDashboardPanelGenericTarget()
}

type DashboardPanelTargetDatasource struct {
	Type string `json:"type"`
	UID  string `json:"uid"`
}

type DashboardPanelGenericTargetImpl[T any] struct {
	CustomJSON customJSON

	DashboardPanelGenericTarget

	Data T
}

type DashboardPanelTargetQuery struct {
	Datasource DashboardPanelTargetDatasource `json:"datasource"`
	RefID      string                         `json:"refId"`
	Hide       bool                           `json:"hide,omitempty"`
}

type DashboardPanelTargetSQL struct {
	DashboardPanelTargetQuery `json:",inline"`

	Format QueryFormat `json:"format"`
	RawSQL string      `json:"rawSql"`
}

type DashboardPanelTargetAthenaSQL struct {
	DashboardPanelTargetQuery `json:",inline"`

	Format         QueryFormatAthenaSQL                        `json:"format"`
	RawSQL         string                                      `json:"rawSQL"`
	ConnectionArgs DashboardPanelTargetAthenaSQLConnectionArgs `json:"connectionArgs"`
}

type DashboardPanelTargetAthenaSQLConnectionArgs struct {
	Region                     string  `json:"region"`
	Catalog                    string  `json:"catalog"`
	Database                   string  `json:"database"`
	ResultReuseEnabled         bool    `json:"resultReuseEnabled"`
	ResultReuseMaxAgeInMinutes float64 `json:"resultReuseMaxAgeInMinutes"`
}

func (d *DashboardPanelTarget) MarshalJSON() ([]byte, error) {
	type plainTarget DashboardPanelTarget
	return marshalResource((*plainTarget)(d), d.CustomJSON)
}

func (d *DashboardPanelGenericTargetImpl[T]) MarshalJSON() ([]byte, error) {
	return marshalResource(&d.Data, d.CustomJSON)
}
