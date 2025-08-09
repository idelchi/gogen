// Package pw provides secure password generation functionality using a diverse
// set of characters including lowercase, uppercase, numbers, and special characters.
package pw

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	// charSetLower defines the set of lowercase ASCII letters.
	charSetLower = "abcdefghijklmnopqrstuvwxyz"

	// charSetUpper defines the set of uppercase ASCII letters.
	charSetUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// charSetNumbers defines the set of decimal digits.
	charSetNumbers = "0123456789"

	// charSetSpecial an opinionated the set of special characters.
	// Avoids characters that cause issues with command-line usage, e.g.
	// no pipes, redirects, globs, quotes, escapes, or command substitutions.
	charSetSpecial = "@#%^_+-=:,."

	// allChars combines all character sets for password generation.
	allChars = charSetLower + charSetUpper + charSetNumbers + charSetSpecial
)

// secureRandomInt generates a cryptographically secure random integer in the range [0, max).
func secureRandomInt(upperBound int) (int, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(int64(upperBound)))
	if err != nil {
		return 0, fmt.Errorf("generating random number: %w", err)
	}

	return int(bigInt.Int64()), nil
}

// Generate creates a password of the specified length using a mix of character classes.
// If requireAll is true, ensures at least one character from each character set.
//
//nolint:gocognit,nestif	// Function complexity is acceptable.
func Generate(length int, requireAll bool) (string, error) {
	if length <= 0 {
		//nolint:err113 // Occasional dynamic errors are fine.
		return "", errors.New("length must be greater than 0")
	}

	if requireAll && length < 4 {
		//nolint:err113 // Occasional dynamic errors are fine.
		return "", errors.New("length must be at least 4 when requiring all character types")
	}

	result := make([]byte, length)

	if requireAll {
		// Place one required character from each set
		charSets := []string{charSetLower, charSetUpper, charSetNumbers, charSetSpecial}

		for index, charset := range charSets {
			idx, err := secureRandomInt(len(charset))
			if err != nil {
				return "", err
			}

			result[index] = charset[idx]
		}

		// Fill remaining positions
		for index := 4; index < length; index++ {
			idx, err := secureRandomInt(len(allChars))
			if err != nil {
				return "", err
			}

			result[index] = allChars[idx]
		}

		// Shuffle to avoid predictable positioning
		for index := len(result) - 1; index > 0; index-- {
			j, err := secureRandomInt(index + 1)
			if err != nil {
				return "", err
			}

			result[index], result[j] = result[j], result[index]
		}
	} else {
		for index := range length {
			idx, err := secureRandomInt(len(allChars))
			if err != nil {
				return "", err
			}

			result[index] = allChars[idx]
		}
	}

	return string(result), nil
}
