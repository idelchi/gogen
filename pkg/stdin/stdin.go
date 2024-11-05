// Package stdin provides simple utilities for reading from stdin.
package stdin

import (
	"fmt"
	"io"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"

	termutil "github.com/andrew-d/go-termutil"
)

// IsPiped checks if something has been piped to stdin.
func IsPiped() bool {
	fi, err := os.Stdin.Stat()

	return (fi.Mode()&os.ModeCharDevice) == 0 && err == nil
}

// Read returns stdin as a string, trimming the trailing newline.
func Read() (string, error) {
	bytes, err := io.ReadAll(os.Stdin)

	return strings.TrimSuffix(string(bytes), "\n"), err
}

func IsPipedPotentially() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	// If it's not a pipe/device (i.e., it's a terminal), return false immediately
	if (fi.Mode() & os.ModeCharDevice) != 0 {
		return false
	}

	// At this point we know it's a pipe, so check if it's /dev/null
	// by attempting a non-blocking read
	buf := make([]byte, 1)
	if n, err := os.Stdin.Read(buf); err != nil || n == 0 {
		return false
	}

	// Got real data, seek back to start
	os.Stdin.Seek(0, 0)
	return true
}

// MaybePiped checks if something has been piped to stdin.
func MaybePiped() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, fmt.Errorf("getting stdin stat: %w", err)
	}

	isPipe := (stat.Mode()&os.ModeNamedPipe) != 0 ||
		(stat.Mode()&(os.ModeCharDevice|os.ModeDir|os.ModeSymlink)) == 0

	return isPipe, nil
}

// MaybePipedTermUtil checks if something has been piped to stdin.
func MaybePipedTermUtil() bool {
	return !termutil.Isatty(os.Stdin.Fd())
}

// MaybePipedIsAtty checks if something has been piped to stdin.
func MaybePipedIsAtty() bool {
	return !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd())
}

// IsPiped checks if data is being piped to stdin, with special handling for
// GitHub Actions and file inputs. It returns:
// - isPiped: true if stdin has data piped to it
// - isNullPipe: true if it's a null device pipe (like in GitHub Actions)
// - err: any error encountered during checking
func IsPipedGithub() (isPiped bool, isNullPipe bool, err error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false, false, err
	}

	// Check if stdin is a pipe/device
	isPiped = (fi.Mode() & os.ModeCharDevice) == 0

	// If we detect a pipe, let's check if it's /dev/null or empty
	if isPiped {
		// Read a single byte to check if it's actually /dev/null or empty
		buf := make([]byte, 1)
		n, err := os.Stdin.Read(buf)

		if err == io.EOF || n == 0 {
			// Empty pipe or /dev/null
			isNullPipe = true
			isPiped = false
		} else if err != nil {
			return false, false, err
		} else {
			// Valid data in pipe - seek back to start
			if _, err := os.Stdin.Seek(0, 0); err != nil {
				return false, false, err
			}
		}
	}

	return isPiped, isNullPipe, nil
}

// IsInputFromStdin determines whether input should be read from stdin based on:
// - Whether stdin is piped
// - Whether a valid input file is specified
// - Command line arguments
func IsInputFromStdin() bool {
	isPiped, isNullPipe, err := IsPipedGithub()
	if err != nil {
		// If we can't determine stdin status, assume no stdin input
		return false
	}

	// Don't read from stdin if:
	// 1. nullInput flag is set
	// 2. It's a null pipe (like in GitHub Actions)
	// 3. No real pipe detected
	// 4. More than one argument
	// 5. One argument that's a valid file
	if isNullPipe || !isPiped {
		return false
	}

	return isPiped
}
