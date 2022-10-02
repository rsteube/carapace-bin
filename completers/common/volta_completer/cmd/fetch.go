package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches a tool to the local machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().BoolP("help", "h", false, "Prints help information")
	fetchCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	fetchCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
