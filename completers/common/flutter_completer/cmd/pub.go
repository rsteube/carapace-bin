package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Commands for managing Flutter packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pubCmd).Standalone()

	pubCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.AddCommand(pubCmd)
}
