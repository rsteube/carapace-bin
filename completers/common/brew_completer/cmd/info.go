package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Display brief statistics for your Homebrew installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("all", false, "Print JSON of all available formulae")
	infoCmd.Flags().Bool("analytics", false, "List global Homebrew analytics data")
	infoCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	infoCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	infoCmd.Flags().String("category", "", "Which type of analytics data to retrieve")
	infoCmd.Flags().String("days", "", "How many days of analytics data to retrieve")
	infoCmd.Flags().BoolP("debug", "d", false, "Display any debugging information")
	infoCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	infoCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	infoCmd.Flags().Bool("github", false, "Open the GitHub source page for formula and cask in a browser")
	infoCmd.Flags().BoolP("help", "h", false, "Show this message")
	infoCmd.Flags().Bool("installed", false, "Print JSON of formulae that are currently installed")
	infoCmd.Flags().Bool("json", false, "Print a JSON representation")
	infoCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	infoCmd.Flags().BoolP("verbose", "v", false, "Show more verbose analytics data for formula")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"category": carapace.ActionValues("install", "install-on-request", "build-error"),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionSearch(installCmd),
	)
}
