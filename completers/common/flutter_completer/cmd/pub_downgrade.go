package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_downgradeCmd = &cobra.Command{
	Use:   "downgrade",
	Short: "Downgrade the current package's dependencies to oldest versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_downgradeCmd).Standalone()

	pub_downgradeCmd.Flags().StringP("directory", "C", "", "Run this in the directory<dir>.")
	pub_downgradeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	pub_downgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_downgradeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_downgradeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pubCmd.AddCommand(pub_downgradeCmd)

	carapace.Gen(pub_downgradeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pub.ActionDependencies().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
