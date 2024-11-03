package commands

import (
	"fmt"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/hash"
	"github.com/spf13/cobra"
)

// NewHashCommand creates the hash subcommand for password hashing operations.
// It handles password hashing with configurable cost and benchmarking.
func NewHashCommand(cfg *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "hash [flags] password",
		Short: "Hash a password",
		Long:  "Hash a password using bcrypt with configurable cost and benchmarking",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cfg.Password.Password = args[0]
			return validate(cfg, &cfg.Password)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.Password.Benchmark {
				hash.Benchmark(cfg.Password.Password)
				return nil
			}

			hash, err := hash.Password(cfg.Password.Password, cfg.Password.Cost)
			if err != nil {
				return fmt.Errorf("generating hash: %w", err)
			}
			fmt.Printf(hash)

			return nil
		},
	}
}
