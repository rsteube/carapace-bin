package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).PositionalCompletion(
		carapace.ActionValues("bash", "zsh", "fish", "powershell"),
	)
}
