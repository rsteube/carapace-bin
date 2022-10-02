package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/restic_completer/cmd/action"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print internal objects to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalCompletion(
		carapace.ActionValues("pack", "blob", "snapshot", "index", "key", "masterkey", "config", "lock"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO complete more objects
			switch c.Args[0] {
			case "snapshot":
				return action.ActionSnapshotIDs(catCmd)
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
