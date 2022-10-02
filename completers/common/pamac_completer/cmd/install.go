package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pamac_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages from repositories, path or url",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("as-deps", false, "mark all packages installed as a dependency")
	installCmd.Flags().Bool("as-explicit", false, "mark all packages explicitly installed")
	installCmd.Flags().BoolP("download-only", "w", false, "download all packages but do not install/upgrade")
	installCmd.Flags().BoolP("dry-run", "d", false, "only print what would be done but do not run the")
	installCmd.Flags().String("ignore", "", "ignore a package upgrade, multiple packages can be")
	installCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	installCmd.Flags().Bool("no-upgrade", false, "do not check for updates")
	installCmd.Flags().String("overwrite", "", "overwrite conflicting files, multiple patterns can be")
	installCmd.Flags().Bool("upgrade", false, "check for updates")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"ignore": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionPackageSearch().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles(".pkg", ".pkg.tar.gz", ".pkg.tar.xz", ".pkg.tar.zst")
			}

			return carapace.Batch(
				action.ActionPackageGroups(),
				action.ActionPackageSearch(),
			).ToA()
		}),
	)
}
