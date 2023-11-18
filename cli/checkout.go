package cli

import (
	"github.com/FollowTheProcess/dev/app"
	"github.com/spf13/cobra"
)

const checkoutLongAbout string = `
The checkout command lets you easily resume work on an existing project,
whether that project is available locally in your configured development directory,
or if it is on GitHub.

If dev finds your project locally, and you have specified an editor in your config file
it will open it for you. If not, it will just give you the path to your project.

If your project is not local, but on GitHub, dev will clone it for you first, then
do the above.
`

// buildCheckoutCmd builds and returns the checkout subcommand.
func buildCheckoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "checkout PROJECT",
		Args:    cobra.ExactArgs(1),
		Short:   "Checkout an existing development project",
		Long:    checkoutLongAbout,
		Example: "$ dev checkout my-project",
		RunE: func(cmd *cobra.Command, args []string) error {
			project := args[0]
			app, err := app.New()
			if err != nil {
				return err
			}
			return app.Checkout(project)
		},
	}

	return cmd
}
