package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Generate tab-completion scripts for your shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "powershell", "elvish"),
		carapace.ActionValues("rustup", "cargo"),
	)
}
