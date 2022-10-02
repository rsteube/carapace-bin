package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uncordonCmd = &cobra.Command{
	Use:   "uncordon",
	Short: "Mark node as schedulable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uncordonCmd).Standalone()

	uncordonCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	uncordonCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on")
	rootCmd.AddCommand(uncordonCmd)

	carapace.Gen(uncordonCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": action.ActionDryRunModes(),
	})

	carapace.Gen(uncordonCmd).PositionalCompletion(
		action.ActionResources("", "nodes"),
	)
}
