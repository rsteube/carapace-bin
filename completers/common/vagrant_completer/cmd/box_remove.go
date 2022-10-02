package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/vagrant_completer/cmd/action"
	"github.com/spf13/cobra"
)

var box_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove a box",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_removeCmd).Standalone()

	box_removeCmd.Flags().Bool("all", false, "Remove all available versions of the box")
	box_removeCmd.Flags().String("box-version", "", "The specific version of the box to remove")
	box_removeCmd.Flags().BoolP("force", "f", false, "Remove without confirmation.")
	box_removeCmd.Flags().String("provider", "", "The specific provider type for the box to remove")
	boxCmd.AddCommand(box_removeCmd)

	carapace.Gen(box_removeCmd).FlagCompletion(carapace.ActionMap{
		"box-version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionBoxVersions(c.Args[0], box_removeCmd.Flag("provider").Value.String())
			}
			return carapace.ActionValues()
		}),
		"provider": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionBoxProviders(c.Args[0], box_removeCmd.Flag("provider").Value.String())
			}
			return action.ActionProviders()
		}),
	})

	carapace.Gen(box_removeCmd).PositionalCompletion(
		action.ActionBoxes(),
	)
}
