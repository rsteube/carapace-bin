package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a dependency from the current package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_removeCmd).Standalone()

	pub_removeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change")
	pub_removeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_removeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_removeCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	pub_removeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pub_removeCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	pubCmd.AddCommand(pub_removeCmd)

	carapace.Gen(pub_removeCmd).PositionalCompletion(
		pub.ActionDependencies(),
	)
}
