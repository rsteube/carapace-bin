package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_globalCmd = &cobra.Command{
	Use:   "global",
	Short: "Work with global packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_globalCmd).Standalone()

	pub_globalCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.AddCommand(pub_globalCmd)
}
