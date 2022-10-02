package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Generates Volta completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().BoolP("force", "f", false, "Write over an existing file, if any.")
	completionsCmd.Flags().BoolP("help", "h", false, "Prints help information")
	completionsCmd.Flags().StringP("output", "o", "", "File to write generated completions to")
	completionsCmd.Flags().Bool("quiet", false, "Prevents unnecessary output")
	completionsCmd.Flags().Bool("verbose", false, "Enables verbose diagnostics")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "powershell", "elvish"),
	)
}
