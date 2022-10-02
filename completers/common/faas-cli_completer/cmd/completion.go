package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates shell auto completion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completionCmd.Flags().String("shell", "", "Outputs shell completion, must be bash or zsh")
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "zsh"),
	})
}
