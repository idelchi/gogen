// pwhash is a simple command-line utility that generates bcrypt
// password hashes.
//
// For further info, run `pwhash -h`.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	cost := 8

	flag.IntVar(&cost, "c", cost, "The cost factor for `bcrypt` hashing")

	flag.Usage = func() {
		fmt.Printf("Usage: %s hash [options] <password>\n", filepath.Base(os.Args[0]))
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	switch flag.NArg() {
	case 0:
		fmt.Fprintf(os.Stderr, "password is required as a positional argument\n")

		os.Exit(1)
	case 1:
	default:
		fmt.Fprintf(os.Stderr, "only one positional argument is allowed\n")

		os.Exit(1)
	}

	password := flag.Arg(0)

	hash, err := HashPassword(password, cost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)

		os.Exit(1)
	}

	fmt.Println(hash)
}

// HashPassword generates a bcrypt hash of the given password using the specified cost.
func HashPassword(password string, cost int) (string, error) {
	if cost < bcrypt.MinCost {
		return "", bcrypt.InvalidCostError(cost)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", fmt.Errorf("error generating bcrypt hash: %w", err)
	}

	return string(bytes), nil
}
