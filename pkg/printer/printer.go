// Package printer provides a simple way to print messages to the standard output and standard error streams.
package printer

import (
	"fmt"
	"os"
)

// Stdoutln prints a message to the standard output stream, appending a newline.
func Stdoutln(format string, args ...any) {
	fmt.Println(fmt.Sprintf(format, args...))
}

// Stderrln prints a message to the standard error stream, appending a newline.
func Stderrln(format string, args ...any) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, args...))
}

// Stdout prints a message to the standard output stream.
func Stdout(format string, args ...any) {
	fmt.Print(fmt.Sprintf(format, args...))
}

// Stderr prints a message to the standard error stream.
func Stderr(format string, args ...any) {
	fmt.Fprint(os.Stderr, fmt.Sprintf(format, args...))
}
