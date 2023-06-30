package tftooltest

import (
	"bytes"
	"encoding/json"
	"runtime/debug"
	"testing"

	"github.com/g7r/ggg/tftool/internal"
)

func Test(t *testing.T, fn func()) map[string]string {
	var (
		oldStdout = internal.Stdout
		oldStderr = internal.Stderr
		oldOSExit = internal.OSExit
	)

	var stdout, stderr bytes.Buffer
	internal.Stdout = &stdout
	internal.Stderr = &stderr
	internal.OSExit = func(code int) {
		t.Fatalf("unexpected call to os.Exit(%d):\n%s", code, debug.Stack())
	}

	defer func() {
		internal.Stdout = oldStdout
		internal.Stderr = oldStderr
		internal.OSExit = oldOSExit
	}()

	runFn(t, fn)

	if stderr.Len() != 0 {
		t.Errorf("stderr should be empty")
	}

	if stdout.Len() == 0 {
		t.Errorf("stdout shouldn't be empty")
	}

	if t.Failed() {
		t.FailNow()
	}

	var dashboards map[string]string
	if err := json.Unmarshal(stdout.Bytes(), &dashboards); err != nil {
		t.Fatalf("failed to unmarshal dashboards: %v", err)
	}

	return dashboards
}

func runFn(t *testing.T, fn func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("dashboard generator shouldn't panic\n%s", debug.Stack())
		}
	}()

	fn()
}
