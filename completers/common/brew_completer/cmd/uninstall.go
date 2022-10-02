package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a formula or cask",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()

	uninstallCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	uninstallCmd.Flags().Bool("casks", false, "Treat all named arguments as casks.")
	uninstallCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	uninstallCmd.Flags().BoolP("force", "f", false, "Delete all installed versions of formula.")
	uninstallCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	uninstallCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae.")
	uninstallCmd.Flags().BoolP("help", "h", false, "Show this message.")
	uninstallCmd.Flags().Bool("ignore-dependencies", false, "Don't fail uninstall , even if formula is a dependency")
	uninstallCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	uninstallCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	uninstallCmd.Flags().Bool("zap", false, "Remove all files associated with a cask.")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(uninstallCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionList(uninstallCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
