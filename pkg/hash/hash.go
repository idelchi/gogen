// Package hash provides functionality for secure password hashing and benchmarking
// using the bcrypt algorithm.
//
// The package offers two main functionalities:
//   - Password hashing with configurable cost factor
//   - Benchmarking tool to measure hashing performance
//
// Example usage:
//
//	// Hash a password with cost of 12
//	hash, err := hash.Password("password", 12)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	// Benchmark hashing performance
//	hash.Benchmark("password")
//
// Note that bcrypt has an upper limit on password length of 72 bytes.
package hash

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Password generates a bcrypt hash of the given password using the specified cost.
// It returns an error if the cost is not within bcrypt's minimum/maximum range.
func Password(password string, cost int) (string, error) {
	switch {
	case cost < bcrypt.MinCost:
		return "", bcrypt.InvalidCostError(cost)
	case cost > bcrypt.MaxCost:
		return "", bcrypt.InvalidCostError(cost)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("error generating bcrypt hash: %w", err)
	}

	return string(bytes), nil
}

// Benchmark prints a table showing the time taken to hash and verify a password
// using bcrypt with different cost factors. It tests all valid cost factors
// from MinCost to MaxCost, measuring both hashing and verification time.
// The output is formatted as a Markdown table.
//
//nolint:forbidigo
func Benchmark(password string) {
	pwd := []byte(password)

	fmt.Println("| Cost Factor  | Estimated Time    |")
	fmt.Println("|--------------|-------------------|")

	for cost := bcrypt.MinCost; cost <= bcrypt.MaxCost; cost++ {
		start := time.Now()

		hash, err := bcrypt.GenerateFromPassword(pwd, cost)
		if err != nil {
			fmt.Printf("| %-12d | Error             |\n", cost)

			continue
		}

		err = bcrypt.CompareHashAndPassword(hash, pwd)
		if err != nil {
			fmt.Printf("error verifying at cost %d: %v\n", cost, err)

			continue
		}

		elapsed := time.Since(start)

		fmt.Printf("| %-12d | %-17s |\n", cost, elapsed)
	}
}
