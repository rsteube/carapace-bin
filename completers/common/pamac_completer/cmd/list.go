package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List packages, groups, repositories or files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("files", "f", false, "list files owned by the given packages")
	listCmd.Flags().BoolP("foreign", "m", false, "list packages that were not found in the")
	listCmd.Flags().BoolP("groups", "g", false, "list all packages that are members of the given")
	listCmd.Flags().BoolP("installed", "i", false, "list installed packages")
	listCmd.Flags().BoolP("orphans", "o", false, "list packages that were installed as dependencies")
	listCmd.Flags().BoolP("quiet", "q", false, "only print names")
	listCmd.Flags().BoolP("repos", "r", false, "list all packages available in the given repos")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if listCmd.Flag("groups").Changed {
			return pacman.ActionPackageGroups().Invoke(c).Filter(c.Args).ToA()
		} else if listCmd.Flag("repos").Changed {
			return pacman.ActionRepositories().Invoke(c).Filter(c.Args).ToA()
		} else if listCmd.Flag("files").Changed {
			return pacman.ActionPackages(pacman.PackageOption{}).Invoke(c).Filter(c.Args).ToA()
		} else {
			return carapace.ActionValues()
		}
	}))
}
