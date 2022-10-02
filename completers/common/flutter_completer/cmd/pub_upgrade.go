package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pub"
	"github.com/spf13/cobra"
)

var pub_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the current package's dependencies to latest versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_upgradeCmd).Standalone()

	pub_upgradeCmd.Flags().StringP("directory", "C", "", "Run this in the directory<dir>.")
	pub_upgradeCmd.Flags().BoolP("dry-run", "n", false, "Report what dependencies would change but don't change any.")
	pub_upgradeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_upgradeCmd.Flags().Bool("major-versions", false, "Upgrades packages to their latest resolvable versions, and updates pubspec.yaml.")
	pub_upgradeCmd.Flags().Bool("no-offline", false, "Do not use cached packages instead of accessing the network.")
	pub_upgradeCmd.Flags().Bool("no-precompile", false, "Do not precompile executables in immediate dependencies.")
	pub_upgradeCmd.Flags().Bool("null-safety", false, "Upgrade constraints in pubspec.yaml to null-safety versions")
	pub_upgradeCmd.Flags().Bool("offline", false, "Use cached packages instead of accessing the network.")
	pub_upgradeCmd.Flags().Bool("precompile", false, "Precompile executables in immediate dependencies.")
	pubCmd.AddCommand(pub_upgradeCmd)

	carapace.Gen(pub_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})

	carapace.Gen(pub_upgradeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pub.ActionDependencies().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
