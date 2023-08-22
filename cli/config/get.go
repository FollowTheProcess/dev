package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAllowedArgs = []string{
	"directory",
	"github.username",
	"github.token",
	"editor.bin",
	"editor.name",
	"editor.open",
}

// buildGetmd builds and returns the config get subcommand.
func buildGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "get KEY",
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short:     "Fetch a specific config value by name.",
		Example:   "$ dev config get github.username",
		ValidArgs: getAllowedArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("config get called with key: %v\n", args[0])
			return nil
		},
	}

	return cmd
}
