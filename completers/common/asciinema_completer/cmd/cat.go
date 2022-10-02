package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print full output of recorded session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
