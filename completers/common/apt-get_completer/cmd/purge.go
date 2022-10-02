package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge is identical to remove except that packages are removed and purged",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	rootCmd.AddCommand(purgeCmd)

	carapace.Gen(purgeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return apt.ActionPackages().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
