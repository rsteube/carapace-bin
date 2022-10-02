package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var commandPromptCmd = &cobra.Command{
	Use:   "command-prompt",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandPromptCmd).Standalone()

	commandPromptCmd.Flags().StringS("I", "I", "", "inputs")
	commandPromptCmd.Flags().StringS("p", "p", "", "prompts")
	commandPromptCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(commandPromptCmd)
}
