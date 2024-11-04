// Package key provides functionality for creating and managing cryptographic keys.
//
// The package supports:
//   - Generating cryptographically secure random keys of arbitrary length
//   - Converting between raw bytes and hexadecimal string representations
//
// Example usage:
//
//	// Generate a new 32-byte key
//	key, err := key.New(32)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Convert to hex string for storage
//	hexKey := key.AsHex()
//
//	// Later, recreate the key from hex
//	restoredKey, err := key.FromHex(hexKey)
//	if err != nil {
//	    log.Fatal(err)
//	}
package key

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

// Key represents a cryptographic key as a byte slice.
type Key []byte

// New creates a new Key of the specified length using cryptographically secure random bytes.
// It returns an error if the random number generator fails.
func New(length int) (Key, error) {
	key := make([]byte, length)

	res, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("error generating random bytes: %w", err)
	}

	if res != length {
		return nil, fmt.Errorf("generated %d bytes instead of requested %d bytes", res, length) //nolint: err113
	}

	return key, nil
}

// FromHex creates a Key by decoding a hexadecimal string.
// It trims any whitespace from the input string before decoding.
// Returns an error if the hex string is invalid.
func FromHex(hexKey string) (Key, error) {
	key, err := hex.DecodeString(strings.TrimSpace(hexKey))
	if err != nil {
		return nil, fmt.Errorf("invalid hex key: %w", err)
	}

	return key, nil
}

// AsHex returns the Key as a lowercase hexadecimal string.
func (k Key) AsHex() string {
	return hex.EncodeToString(k)
}
