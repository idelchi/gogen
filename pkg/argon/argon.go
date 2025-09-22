// Package argon provides functionality for secure password hashing using Argon2id.
package argon

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

// Password generates an Argon2id hash of the provided password using default parameters.
// Returns a string in the format: $argon2id$v=19$m=65536,t=3,p=2$<salt>$<hash>.
func Password(password string) (string, error) {
	defaults := argon2id.DefaultParams

	defaults.Iterations = 3
	defaults.Parallelism = 4

	hash, err := argon2id.CreateHash(password, defaults)
	if err != nil {
		return "", fmt.Errorf("creating hash: %w", err)
	}

	return hash, nil
}
