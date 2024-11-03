package key

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

type Key []byte

func New(length int) (Key, error) {
	key, err := generate(length)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func FromHex(hexKey string) (Key, error) {
	key, err := hex.DecodeString(strings.TrimSpace(string(hexKey)))
	if err != nil {
		return nil, fmt.Errorf("invalid hex key: %w", err)
	}

	return key, nil
}

func (k *Key) AsHex() string {
	return hex.EncodeToString((*k))
}

// generate generates a 32-byte key for AES-256 encryption.
func generate(length int) ([]byte, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("error generating random bytes: %w", err)
	}
	return key, nil
}
