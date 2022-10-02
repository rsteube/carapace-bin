package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open a formula or cask in the editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	editCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	editCmd.Flags().Bool("casks", false, "Treat all named arguments as casks.")
	editCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	editCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	editCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae.")
	editCmd.Flags().BoolP("help", "h", false, "Show this message.")
	editCmd.Flags().Bool("print-path", false, "Print the file path to be edited")
	editCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	editCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).PositionalAnyCompletion(
		action.ActionList(editCmd),
	)
}
