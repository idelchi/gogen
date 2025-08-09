package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/cobraext"
	"github.com/idelchi/gogen/pkg/key"
)

// NewKeyCommand creates the key generation subcommand.
// It handles generating cryptographic keys of specified length.
//
//nolint:forbidigo	// Command prints out to the console.
func NewKeyCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Generate a cryptographic key",
		Long:  "Generate a cryptographic key of specified length",
		Args:  cobra.NoArgs,
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return cobraext.Validate(cfg, &cfg.Generate)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			key, err := key.New(cfg.Generate.Length)
			if err != nil {
				return fmt.Errorf("generating key: %w", err)
			}

			fmt.Print(key.AsHex())

			return nil
		},
	}

	const length = 32

	cmd.Flags().IntP("length", "l", length, "Length of the key to generate")

	return cmd
}
