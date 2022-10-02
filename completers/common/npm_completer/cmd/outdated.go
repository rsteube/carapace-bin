package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Check for outdated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()
	outdatedCmd.Flags().BoolP("all", "a", false, "show all")
	outdatedCmd.Flags().BoolP("global", "g", false, "operate globally")
	outdatedCmd.Flags().Bool("json", false, "output as json")
	outdatedCmd.Flags().BoolP("long", "l", false, "show extended information")
	outdatedCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	addWorkspaceFlags(outdatedCmd)

	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		action.ActionPackages(outdatedCmd),
	)
}
