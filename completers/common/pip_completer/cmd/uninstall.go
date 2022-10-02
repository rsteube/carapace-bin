package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().StringP("requirement", "r", "", "Uninstall all the packages listed in the given requirements file.")
	uninstallCmd.Flags().BoolP("yes", "y", false, "Don't ask for confirmation of uninstall deletions.")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).FlagCompletion(carapace.ActionMap{
		"requirement": carapace.ActionFiles(),
	})

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if uninstallCmd.Flag("requirement").Changed {
				return carapace.ActionValues()
			} else {
				return pip.ActionInstalledPackages()
			}
		}),
	)
}
