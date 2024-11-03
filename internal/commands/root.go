package commands

import (
	"fmt"
	"strings"

	"github.com/idelchi/gogen/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRootCommand creates the root command with common configuration.
// It sets up environment variable binding and flag handling.
func NewRootCommand(_ *config.Config) *cobra.Command {
	root := &cobra.Command{
		SilenceUsage:     true,
		SilenceErrors:    true,
		Use:              "gogen [flags] command [flags]",
		Short:            "Generate cryptographic keys and password hashes",
		Long:             "gogen is a tool for generating cryptographic keys and password hashes.",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetEnvPrefix(cmd.Root().Name())
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
			viper.AutomaticEnv()

			if err := viper.BindPFlags(cmd.Root().Flags()); err != nil {
				return fmt.Errorf("failed to bind flags: %w", err)
			}

			if err := viper.BindPFlags(cmd.Flags()); err != nil {
				return fmt.Errorf("failed to bind persistent flags: %w", err)
			}

			return nil
		},
	}

	root.CompletionOptions.DisableDefaultCmd = true
	root.SetVersionTemplate("{{ .Version }}\n")

	return root
}
