package commands

import (
	"fmt"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/key"
	"github.com/spf13/cobra"
)

// NewKeyCommand creates the key generation subcommand.
// It handles generating cryptographic keys of specified length.
func NewKeyCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Generate a cryptographic key",
		Long:  "Generate a cryptographic key of specified length",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validate(cfg, &cfg.Generate)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			key, err := key.New(cfg.Generate.Length)
			if err != nil {
				return fmt.Errorf("generating key: %w", err)
			}
			fmt.Printf(key.AsHex())

			return nil
		},
	}
	return cmd
}
