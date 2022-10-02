package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Locates the actual binary that will be called by Volta",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().BoolP("help", "h", false, "Prints help information")
	whichCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	whichCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).PositionalCompletion(
		os.ActionPathExecutables(),
	)
}
