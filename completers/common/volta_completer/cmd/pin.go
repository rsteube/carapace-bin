package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Pins your project's runtime or package manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinCmd).Standalone()

	pinCmd.Flags().BoolP("help", "h", false, "Prints help information")
	pinCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	pinCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(pinCmd)

	carapace.Gen(pinCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
