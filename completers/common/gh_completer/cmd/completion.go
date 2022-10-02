package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()
	completionCmd.Flags().StringP("shell", "s", "", "Shell type: {bash|zsh|fish|powershell}")
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "fish", "powershell", "zsh"),
	})
}
