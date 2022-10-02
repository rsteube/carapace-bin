package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pamac_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("dry-run", "d", false, "only print what would be done but do not run the transaction")
	removeCmd.Flags().Bool("no-confirm", false, "bypass any and all confirmation messages")
	removeCmd.Flags().Bool("no-orphans", false, "do not remove dependencies that are not required by other")
	removeCmd.Flags().BoolP("no-save", "n", false, "ignore files backup")
	removeCmd.Flags().BoolP("orphans", "o", false, "remove dependencies that are not required by other packages,")
	removeCmd.Flags().BoolP("unneeded", "u", false, "remove packages only if they are not required by any other")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return action.ActionInstalledPackages(true).Invoke(c).Filter(c.Args).ToA()
	}))
}
