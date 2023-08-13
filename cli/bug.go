package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildBugCmd builds and returns the bug subcommand.
func buildBugCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "bug",
		Args:    cobra.NoArgs,
		Short:   "File an issue about dev.",
		Example: "$ dev bug",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("bug called")
			return nil
		},
	}

	return cmd
}
