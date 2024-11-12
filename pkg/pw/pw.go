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

	// charSetSpecial defines the set of special characters.
	charSetSpecial = "!@#$%^&*()_+-=[]{}|;:,.<>?"

	// allChars combines all character sets for password generation.
	allChars = charSetLower + charSetUpper + charSetNumbers + charSetSpecial
)

// secureRandomInt generates a cryptographically secure random integer in the range [0, max).
// It uses crypto/rand to ensure high-quality randomness suitable for security-sensitive operations.
func secureRandomInt(upperBound int) (int, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(int64(upperBound)))
	if err != nil {
		return 0, fmt.Errorf("generating random number: %w", err)
	}

	return int(bigInt.Int64()), nil
}

// Generate creates a password of the specified length using a mix of character classes.
// The generated password uses a diverse character set including lowercase, uppercase,
// numbers, and special characters, selected using cryptographically secure random numbers.
func Generate(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0") //nolint:err113
	}

	// Initialize password builder
	result := make([]byte, length)

	// Fill positions with random characters from all classes
	for index := range length {
		idx, err := secureRandomInt(len(allChars))
		if err != nil {
			return "", err
		}

		result[index] = allChars[idx]
	}

	return string(result), nil
}
