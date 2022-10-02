package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pub_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current package's dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_getCmd).Standalone()

	pub_getCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	pub_getCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_getCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_getCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	pub_getCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pub_getCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	pubCmd.AddCommand(pub_getCmd)
}
