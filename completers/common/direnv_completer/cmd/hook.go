package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Used to setup the shell hook",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hookCmd).Standalone()

	rootCmd.AddCommand(hookCmd)

	carapace.Gen(hookCmd).PositionalCompletion(
		carapace.ActionValues("bash", "elvish", "fish", "tcsh", "zsh"),
	)
}
