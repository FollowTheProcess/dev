// Package cli implements dev's command line interface.
package cli

import "github.com/spf13/cobra"

// These are all set at compile time.
var (
	version   = "dev"
	commit    = ""
	buildDate = ""
	builtBy   = ""
)

// Build builds and returns the dev CLI.
func Build() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "dev COMMAND [FLAGS]",
		Version:       version,
		Args:          cobra.NoArgs,
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "The all in one developer toolkit 🛠️",
	}

	// Set our custom version and usage templates
	cmd.SetUsageTemplate(usageTemplate)
	cmd.SetVersionTemplate(versionTemplate)

	// Attach the subcommands
	cmd.AddCommand(buildBugCmd())

	return cmd
}
