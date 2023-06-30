package internal

import (
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr
var OSExit = os.Exit
