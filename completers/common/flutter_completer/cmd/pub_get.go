package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get packages in a Flutter project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_getCmd).Standalone()

	pub_getCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_getCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pubCmd.AddCommand(pub_getCmd)

	carapace.Gen(pub_getCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
