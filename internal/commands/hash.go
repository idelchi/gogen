package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/cobraext"
	"github.com/idelchi/gogen/pkg/hash"
	"github.com/idelchi/gogen/pkg/printer"
	"github.com/idelchi/gogen/pkg/stdin"
)

// NewHashCommand creates the hash subcommand for password hashing operations.
// It handles password hashing with configurable cost and benchmarking.
func NewHashCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hash [flags] password",
		Short: "Hash a password",
		Long:  "Hash a password using bcrypt with configurable cost and benchmarking. Can read password from stdin.",
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(_ *cobra.Command, args []string) error {
			isPiped := stdin.IsPiped()

			switch {
			case len(args) > 0:
				// Prioritize argument if it exists, regardless of stdin
				cfg.Password.Password = args[0]
				if isPiped {
					printer.Stderrln("reading password from argument, ignoring stdin")
				}
			case isPiped:
				// No arg but stdin is piped
				password, err := stdin.Read()
				if err != nil {
					return fmt.Errorf("reading password from stdin: %w", err)
				}
				cfg.Password.Password = password
			}

			return cobraext.Validate(cfg, &cfg.Password)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			if cfg.Password.Benchmark {
				hash.Benchmark(cfg.Password.Password)

				return nil
			}

			hash, err := hash.Password(cfg.Password.Password, cfg.Password.Cost)
			if err != nil {
				return fmt.Errorf("generating hash: %w", err)
			}

			fmt.Print(hash) //nolint: forbidigo

			return nil
		},
	}

	const cost = 12

	cmd.Flags().IntP("cost", "c", cost, "Cost of the password hash (4-31)")
	cmd.Flags().BoolP("benchmark", "b", false, "Run a benchmark on the password hash")

	return cmd
}
