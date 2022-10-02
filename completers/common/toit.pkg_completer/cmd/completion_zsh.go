package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate the autocompletion script for zsh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_zshCmd).Standalone()
	completion_zshCmd.Flags().Bool("no-descriptions", false, "disable completion descriptions")
	completionCmd.AddCommand(completion_zshCmd)
}
