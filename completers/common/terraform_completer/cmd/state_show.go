package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var state_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a resource in the state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(state_showCmd).Standalone()
	state_showCmd.Flags().StringS("state", "state", "", "Path to a Terraform state file")
	stateCmd.AddCommand(state_showCmd)

	state_showCmd.Flag("state").NoOptDefVal = " "

	carapace.Gen(state_showCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionFiles(),
	})
}
