package config

import (
	"fmt"

	"github.com/FollowTheProcess/dev/app"
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
	app, err := app.New()
	cmd := &cobra.Command{
		Use:       "get KEY",
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Short:     "Fetch a specific config value by name.",
		Example:   "$ dev config get github.username",
		ValidArgs: getAllowedArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err != nil {
				return err
			}
			fmt.Printf("config get called with key: %v\n", args[0])
			fmt.Printf("%+v\n", app.Config())
			return nil
		},
	}

	return cmd
}
