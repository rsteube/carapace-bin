package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cordonCmd = &cobra.Command{
	Use:   "cordon",
	Short: "Mark node as unschedulable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cordonCmd).Standalone()

	cordonCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent,")
	cordonCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on")
	rootCmd.AddCommand(cordonCmd)

	carapace.Gen(cordonCmd).FlagCompletion(carapace.ActionMap{
		"dry-run": action.ActionDryRunModes(),
	})

	carapace.Gen(cordonCmd).PositionalCompletion(
		action.ActionResources("", "nodes"),
	)
}
