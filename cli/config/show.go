package config

import (
	"github.com/FollowTheProcess/dev/app"
	"github.com/spf13/cobra"
)

// buildShowCmd builds and returns the config show subcommand.
func buildShowCmd() *cobra.Command {
	app, err := app.New()
	cmd := &cobra.Command{
		Use:     "show",
		Args:    cobra.NoArgs,
		Short:   "Display the current config.",
		Example: "$ dev config show",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err != nil {
				return err
			}
			return app.Config().Show(app.Stdout)
		},
	}

	return cmd
}
