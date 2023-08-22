package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildEditCmd builds and returns the config edit subcommand.
func buildEditCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit",
		Args:    cobra.NoArgs,
		Short:   "Open the config file for editing.",
		Example: "$ dev config edit",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("config edit called")
			return nil
		},
	}

	return cmd
}
