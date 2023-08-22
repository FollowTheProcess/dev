package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildShowCmd builds and returns the config show subcommand.
func buildShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "show",
		Args:    cobra.NoArgs,
		Short:   "Display the current config.",
		Example: "$ dev config show",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("config show called")
			return nil
		},
	}

	return cmd
}
