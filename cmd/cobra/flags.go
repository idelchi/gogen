package main

import (
	"github.com/idelchi/gogen/internal/commands"
	"github.com/idelchi/gogen/internal/config"
	"github.com/spf13/cobra"
)

// flags creates and configures the command-line interface.
// It returns the root command with all subcommands and flags configured.
func flags() *cobra.Command {
	cfg := &config.Config{}
	root := commands.NewRootCommand(cfg)
	root.Version = version

	root.Flags().BoolP("show", "s", false, "Show the configuration and exit")

	password := commands.NewHashCommand(cfg)
	generate := commands.NewKeyCommand(cfg)

	root.AddCommand(password, generate)

	generate.Flags().IntP("length", "l", 32, "Length of the key to generate")
	password.Flags().IntP("cost", "c", 12, "Cost of the password hash (4-31)")
	password.Flags().BoolP("benchmark", "b", false, "Run a benchmark on the password hash")

	root.CompletionOptions.DisableDefaultCmd = true
	root.SetVersionTemplate("{{ .Version }}\n")
	root.Flags().SortFlags = false

	return root
}
