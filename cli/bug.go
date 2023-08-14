package cli

import (
	"github.com/FollowTheProcess/msg"
	"github.com/cli/browser"
	"github.com/spf13/cobra"
)

const bugURL = "https://github.com/FollowTheProcess/dev/issues/new/choose"

// buildBugCmd builds and returns the bug subcommand.
func buildBugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bug",
		Args:    cobra.NoArgs,
		Short:   "File an issue about dev.",
		Long:    "The bug command will open your default browser on dev's issue page.",
		Example: "$ dev bug",
		RunE: func(cmd *cobra.Command, args []string) error {
			msg.Info("Opening %s in your browser", bugURL)
			return browser.OpenURL(bugURL)
		},
	}

	return cmd
}
