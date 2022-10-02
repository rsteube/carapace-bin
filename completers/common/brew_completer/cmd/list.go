package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all installed formulae and casks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("1", "1", false, "Force output to be one entry per line")
	listCmd.Flags().Bool("cask", false, "List only casks, or treat all named arguments as casks.")
	listCmd.Flags().Bool("casks", false, "List only casks, or treat all named arguments as casks.")
	listCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	listCmd.Flags().Bool("formula", false, "List only formulae, or treat all named arguments as formulae.")
	listCmd.Flags().Bool("formulae", false, "List only formulae, or treat all named arguments as formulae.")
	listCmd.Flags().Bool("full-name", false, "Print formulae with fully-qualified names.")
	listCmd.Flags().BoolP("help", "h", false, "Show this message.")
	listCmd.Flags().BoolS("l", "l", false, "List formulae and/or casks in long format.")
	listCmd.Flags().Bool("multiple", false, "Only show formulae with multiple versions installed")
	listCmd.Flags().Bool("pinned", false, "List only pinned formulae")
	listCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	listCmd.Flags().BoolS("r", "r", false, "Reverse the order of the formulae")
	listCmd.Flags().BoolS("t", "t", false, "Sort formulae and/or casks by time modified")
	listCmd.Flags().Bool("unbrewed", false, "List files in Homebrew's prefix not installed by Homebrew.")
	listCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	listCmd.Flags().Bool("versions", false, "Show the version number for installed formulae")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionList(listCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
