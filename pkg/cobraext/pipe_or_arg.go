package cobraext

import (
	"fmt"

	"github.com/idelchi/gogen/pkg/printer"
	"github.com/idelchi/gogen/pkg/stdin"
)

// PipeOrArg reads a password from either the first argument or stdin.
func PipeOrArg(args []string) (string, error) {
	isPiped := stdin.IsPiped()

	switch {
	case len(args) > 0:
		// Prioritize argument if it exists, regardless of stdin
		if isPiped {
			printer.Stderrln("reading password from argument, ignoring stdin")
		}

		return args[0], nil
	case isPiped:
		// No arg but stdin is piped
		arg, err := stdin.Read()
		if err != nil {
			return "", fmt.Errorf("reading password from stdin: %w", err)
		}

		return arg, nil
	}

	return "", nil
}
