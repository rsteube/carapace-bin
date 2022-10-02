package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_global_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List globally activated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_global_listCmd).Standalone()

	pub_global_listCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_globalCmd.AddCommand(pub_global_listCmd)
}
