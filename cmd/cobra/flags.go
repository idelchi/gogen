package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/idelchi/godyl/pkg/pretty"
	"github.com/idelchi/gogen/pkg/hash"
	"github.com/idelchi/gogen/pkg/key"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func flags() *cobra.Command {
	cfg := &Config{}
	root := newRootCmd(cfg)

	root.Flags().BoolP("show", "s", false, "Show the configuration and exit")

	password := newPasswordCmd(cfg)
	generate := newGenerateCmd(cfg)

	root.AddCommand(password, generate)

	generate.Flags().IntP("length", "l", 32, "Length of the key to generate")
	password.Flags().IntP("cost", "c", 12, "Cost of the password hash (4-31)")
	password.Flags().BoolP("benchmark", "b", false, "Run a benchmark on the password hash")

	root.CompletionOptions.DisableDefaultCmd = true
	root.SetVersionTemplate("{{ .Version }}\n")
	root.Flags().SortFlags = false

	return root
}

func validate(cfg *Config) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshalling config: %w", err)
	}

	if cfg.Show {
		pretty.PrintJSONMasked(cfg)

		os.Exit(0)

		return nil
	}

	if err := cfg.Validate(); err != nil {
		return fmt.Errorf("validating config: %w\nSee --help for more info on usage.", err)
	}

	return nil
}

func newRootCmd(_ *Config) *cobra.Command {
	root := &cobra.Command{
		// SilenceUsage: true,
		// SilenceErrors:    true,
		Version:          version,
		Use:              "gogen [flags] command [flags]",
		Short:            "",
		Long:             "g",
		TraverseChildren: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetEnvPrefix(cmd.Root().Name())
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
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

func newPasswordCmd(cfg *Config) *cobra.Command {
	return &cobra.Command{
		Use:     "password [flags] password",
		Aliases: []string{"pwd"},
		Short:   "",
		Long:    "",
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cfg.Password.Password = args[0]

			if err := validate(cfg); err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.Password.Benchmark {
				hash.Benchmark(cfg.Password.Password, 4, 31)

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

func newGenerateCmd(cfg *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validate(cfg); err != nil {
				return err
			}

			return nil
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

func newJWTCmd(cfg *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jwt",
		Short: "",
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validate(cfg); err != nil {
				return err
			}

			return nil
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
