package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/cobraext"
	"github.com/idelchi/gogen/pkg/pw"
)

// NewPasswordCommand creates the key generation subcommand.
// It handles generating cryptographic keys of specified length.
//
//nolint:forbidigo	// Command prints out to the console.
func NewPasswordCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "password",
		Short:   "Generate a password",
		Long:    "Generate a password of specified length",
		Aliases: []string{"pw"},
		Args:    cobra.NoArgs,
		PreRunE: func(_ *cobra.Command, _ []string) error {
			return cobraext.Validate(cfg, &cfg.Password)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			password, err := pw.Generate(cfg.Password.Length, true)
			if err != nil {
				return fmt.Errorf("generating password: %w", err)
			}

			fmt.Print(password)

			return nil
		},
	}

	const length = 16

	cmd.Flags().IntP("length", "l", length, "Length of the password to generate")

	return cmd
}
