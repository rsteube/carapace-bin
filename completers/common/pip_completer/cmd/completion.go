package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "A helper command to be used for command completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().BoolP("bash", "b", false, "Emit completion code for bash")
	completionCmd.Flags().BoolP("fish", "f", false, "Emit completion code for fish")
	completionCmd.Flags().BoolP("zsh", "z", false, "Emit completion code for zsh")
	rootCmd.AddCommand(completionCmd)
}
