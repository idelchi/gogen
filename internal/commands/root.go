package commands

import (
	"github.com/spf13/cobra"

	"github.com/idelchi/gogen/internal/config"
	"github.com/idelchi/gogen/pkg/cobraext"
)

// NewRootCommand creates the root command with common configuration.
// It sets up environment variable binding and flag handling.
func NewRootCommand(cfg *config.Config, version string) *cobra.Command {
	root := cobraext.NewDefaultRootCommand(version)

	root.Use = "gogen [flags] command [flags]"
	root.Short = "Generate cryptographic keys and password hashes"
	root.Long = "gogen is a tool for generating cryptographic keys and password hashes."

	root.Flags().BoolP("show", "s", false, "Show the configuration and exit")
	root.AddCommand(NewHashCommand(cfg), NewKeyCommand(cfg))

	return root
}
