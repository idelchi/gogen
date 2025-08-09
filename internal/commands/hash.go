package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/argon"
	"github.com/idelchi/gogen/pkg/cobraext"
	"github.com/idelchi/gogen/pkg/hash"
)

// NewHashCommand creates the hash subcommand for password hashing operations.
// It handles password hashing with configurable cost and benchmarking.
//
//nolint:forbidigo	// Command prints out to the console.
func NewHashCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash [flags] [password|STDIN]",
		Short: "Hash a password",
		Long:  "Hash a password using bcrypt with configurable cost and benchmarking.",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(_ *cobra.Command, args []string) error {
			arg, err := cobraext.PipeOrArg(args)
			if err != nil {
				return fmt.Errorf("reading password: %w", err)
			}

			cfg.Hash.Password = arg

			return cobraext.Validate(cfg, &cfg.Hash)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cfg.Hash.Type == "argon" {
				if cfg.Hash.Benchmark {
					return fmt.Errorf("%w: argon does not support benchmarking", config.ErrUsage)
				}

				if cmd.Flags().Lookup("cost").Changed {
					return fmt.Errorf("%w: argon does not support custom cost", config.ErrUsage)
				}
			}

			if cfg.Hash.Benchmark {
				hash.Benchmark(cfg.Hash.Password)

				return nil
			}

			var hashedPassword string
			var err error

			switch cfg.Hash.Type {
			case "bcrypt":
				hashedPassword, err = hash.Password(cfg.Hash.Password, cfg.Hash.Cost)
			case "argon":
				hashedPassword, err = argon.Password(cfg.Hash.Password)
			default:
				return fmt.Errorf("%w: invalid hash type", config.ErrUsage)
			}

			if err != nil {
				return fmt.Errorf("generating hash: %w", err)
			}

			fmt.Print(hashedPassword)

			return nil
		},
	}

	const cost = 12

	cmd.Flags().IntP("cost", "c", cost, "Cost of the password hash (4-31)")
	cmd.Flags().BoolP("benchmark", "b", false, "Run a benchmark on the password hash")
	cmd.Flags().StringP("type", "t", "bcrypt", "Hashing algorithm to use (bcrypt, argon2id)")

	return cmd
}
