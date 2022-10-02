package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stack_historyCmd = &cobra.Command{
	Use:     "history",
	Short:   "[PREVIEW] Display history for a stack",
	Aliases: []string{"hist"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_historyCmd).Standalone()
	stack_historyCmd.PersistentFlags().Bool("full-dates", false, "Show full dates, instead of relative dates")
	stack_historyCmd.PersistentFlags().BoolP("json", "j", false, "Emit output as JSON")
	stack_historyCmd.PersistentFlags().Int("page", 1, "Used with 'page-size' to paginate results")
	stack_historyCmd.PersistentFlags().Int("page-size", 10, "Used with 'page' to control number of results returned")
	stack_historyCmd.Flags().Bool("show-secrets", false, "Show secret values when listing config instead of displaying blinded values")
	stackCmd.AddCommand(stack_historyCmd)
}
