// Package printer provides a simple way to print messages to the standard output and standard error streams.
//
//nolint:forbidigo	// Package prints out to console.
package printer

import (
	"fmt"
	"os"
)

// Stdoutln prints a message to the standard output stream, appending a newline.
func Stdoutln(format string, args ...any) {
	Stdout(format+"\n", args...)
}

// Stderrln prints a message to the standard error stream, appending a newline.
func Stderrln(format string, args ...any) {
	Stderr(format+"\n", args...)
}

// Stdout prints a message to the standard output stream.
func Stdout(format string, args ...any) {
	fmt.Printf(format, args...)
}

// Stderr prints a message to the standard error stream.
func Stderr(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format, args...)
}
