package tftool

import (
	"encoding/json"
	"fmt"

	"github.com/g7r/ggg/schema"
)

func MarshalDashboard(dashboard *schema.Dashboard) string {
	dashboardJSON, err := json.Marshal(dashboard)
	if err != nil {
		panic(err.Error())
	}

	return string(dashboardJSON)
}

func Main(dashboard *schema.Dashboard) {
	var output struct {
		Dashboard string `json:"dashboard"`
	}

	output.Dashboard = MarshalDashboard(dashboard)
	outputJSON, err := json.Marshal(output)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(outputJSON))
}
