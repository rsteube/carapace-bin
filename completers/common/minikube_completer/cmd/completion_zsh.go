package cmd

import (
	"github.com/spf13/cobra"
)

var completion_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "zsh completion.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	completionCmd.AddCommand(completion_zshCmd)
}
