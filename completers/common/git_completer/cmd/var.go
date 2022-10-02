package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var varCmd = &cobra.Command{
	Use:   "var",
	Short: "Show a Git logical variable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(varCmd).Standalone()
	varCmd.Flags().BoolS("l", "l", false, "Cause the logical variables to be listed")

	rootCmd.AddCommand(varCmd)

	carapace.Gen(varCmd).PositionalCompletion(
		git.ActionVariables(),
	)
}
