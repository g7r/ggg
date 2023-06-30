package tftool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime/debug"

	"github.com/g7r/ggg/schema"
	"github.com/g7r/ggg/tftool/internal"
)

func HandleFatalError(prefix string, err error) {
	_, _ = fmt.Fprintf(internal.Stderr, "ERROR: %s: %s\n%s", prefix, err, debug.Stack())
	internal.OSExit(1)
}

func Main(dashboards map[string]*schema.Dashboard) {
	output := map[string]string{}
	for name, dashboard := range dashboards {
		dashboardJSON, err := json.Marshal(dashboard)
		if err != nil {
			HandleFatalError(fmt.Sprintf("failed to marshal dashboard `%s` to JSON", name), err)
		}

		output[name] = string(dashboardJSON)
	}

	outputJSON, err := json.Marshal(output)
	if err != nil {
		HandleFatalError("failed to marshal output to JSON", err)
	}

	_, _ = fmt.Fprintln(internal.Stdout, string(outputJSON))
}

func ParseArgs(args interface{}) {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		HandleFatalError("failed to read stdin", err)
	}

	inputBytes = bytes.TrimSpace(inputBytes)
	if len(inputBytes) == 0 {
		return
	}

	if err := json.Unmarshal(inputBytes, args); err != nil {
		HandleFatalError("failed to unmarshal input JSON (`args` should be compatible with map[string]string)", err)
	}
}
