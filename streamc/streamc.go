// Provides colorable stdout and stderr streams
package streamc

import (
	"github.com/mattn/go-colorable"
)

// Colorable-aware stdout, similar to os.Stdout
var Stdout = colorable.NewColorableStdout()

// Colorable-aware stderr, similar to os.Stderr
var Stderr = colorable.NewColorableStderr()
