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

		// Test verification
		err = bcrypt.CompareHashAndPassword(hash, pwd)
		if err != nil {
			fmt.Printf("error verifying at cost %d: %v\n", cost, err)

			continue
		}

		elapsed := time.Since(start)

		fmt.Printf("| %-12d | %-17s |\n", cost, elapsed)
	}
}
