package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/idelchi/gogen/internal/config"
)

// NewRootCommand creates the root command with common configuration.
// It sets up environment variable binding and flag handling.
func NewRootCommand(cfg *config.Config, version string) *cobra.Command {
	root := &cobra.Command{
		Version:          version,
		SilenceUsage:     true,
		SilenceErrors:    true,
		Use:              "gogen [flags] command [flags]",
		Short:            "Generate cryptographic keys and password hashes",
		Long:             "gogen is a tool for generating cryptographic keys and password hashes.",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
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

	root.Flags().BoolP("show", "s", false, "Show the configuration and exit")
	root.AddCommand(NewHashCommand(cfg), NewKeyCommand(cfg))

	root.CompletionOptions.DisableDefaultCmd = true
	root.Flags().SortFlags = false

	root.SetVersionTemplate("{{ .Version }}\n")

	return root
}
