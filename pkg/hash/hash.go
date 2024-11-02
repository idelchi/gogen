package hash

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Password generates a bcrypt hash of the given password using the specified cost.
func Password(password string, cost int) (string, error) {
	if cost < bcrypt.MinCost {
		return "", bcrypt.InvalidCostError(cost)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("error generating bcrypt hash: %w", err)
	}

	return string(bytes), nil
}

func Benchmark(password string, minCost, maxCost int) {
	fmt.Printf("Benchmarking bcrypt costs from %d to %d\n", minCost, maxCost)
	fmt.Println("----------------------------------------")
	fmt.Printf("%-12s | %-15s\n", "Cost Factor", "Duration")
	fmt.Println("----------------------------------------")

	for cost := minCost; cost <= maxCost; cost++ {
		start := time.Now()

		// Generate hash
		hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			fmt.Printf("error at cost %d: %v\n", cost, err)

			continue
		}

		// Test verification
		err = bcrypt.CompareHashAndPassword(hash, []byte(password))
		if err != nil {
			fmt.Printf("error verifying at cost %d: %v\n", cost, err)

			continue
		}

		duration := time.Since(start)
		fmt.Printf("%-12d | %v\n", cost, duration)
	}
}
