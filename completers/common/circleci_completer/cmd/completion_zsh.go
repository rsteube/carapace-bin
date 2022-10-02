package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completion_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Generate zsh completion scripts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completion_zshCmd).Standalone()
	completionCmd.AddCommand(completion_zshCmd)
}
