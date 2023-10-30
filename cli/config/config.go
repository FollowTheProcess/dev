// Package cli/config implements the config subcommand group.
package config

import "github.com/spf13/cobra"

// BuildConfigCmd builds and returns the config subcommand group.
func BuildConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config SUBCOMMAND [FLAGS]",
		Args:  cobra.NoArgs,
		Short: "Interact with dev's configuration.",
	}

	cmd.AddCommand(
		buildShowCmd(),
		buildEditCmd(),
	)

	return cmd
}
