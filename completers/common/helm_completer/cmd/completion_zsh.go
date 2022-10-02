package cmd

import (
	"github.com/spf13/cobra"
)

var completion_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "generate autocompletion script for zsh",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completion_zshCmd.Flags().Bool("no-descriptions", false, "disable completion descriptions")
	completionCmd.AddCommand(completion_zshCmd)
}
