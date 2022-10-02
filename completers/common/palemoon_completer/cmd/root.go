package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "palemoon",
	Short: "Pale Moon browser",
	Long:  "https://www.palemoon.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("P", "P", "", "Start with <profile>.")
	rootCmd.Flags().Bool("ProfileManager", false, "Start with ProfileManager.")
	rootCmd.Flags().String("UILocale", "", "Start with <locale> resources as UI Locale.")
	rootCmd.Flags().Bool("browser", false, "Open a browser window.")
	rootCmd.Flags().Bool("browserconsole", false, "Open the Browser Console.")
	rootCmd.Flags().Bool("devtools", false, "Open DevTools on initial load.")
	rootCmd.Flags().String("display", "", "X display to use.")
	rootCmd.Flags().Bool("g-fatal-warnings", false, "Make all warnings fatal.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this message.")
	rootCmd.Flags().Bool("jsconsole", false, "Open the Error console.")
	rootCmd.Flags().Bool("jsdebugger", false, "Open the Browser Toolbox.")
	rootCmd.Flags().Bool("new-instance", false, "Open new instance")
	rootCmd.Flags().String("new-tab", "", "Open <url> in a new tab.")
	rootCmd.Flags().String("new-window", "", "Open <url> in a new window.")
	rootCmd.Flags().Bool("no-remote", false, "Do not accept or send remote commands;")
	rootCmd.Flags().Bool("preferences", false, "Open Preferences dialog.")
	rootCmd.Flags().String("private-window", "", "Open <url> in a new private window.")
	rootCmd.Flags().String("profile", "", "Start with profile at <path>.")
	rootCmd.Flags().String("recording", "", "Record drawing for a given URL.")
	rootCmd.Flags().String("recording-output", "", "Specify destination file for a drawing recording.")
	rootCmd.Flags().Bool("safe-mode", false, "Disables extensions and themes for this session.")
	rootCmd.Flags().String("search", "", "Search <term> with your default search engine.")
	rootCmd.Flags().Bool("setDefaultBrowser", false, "Set this app as the default browser.")
	rootCmd.Flags().String("start-debugger-server", "", "Start the debugger server on a TCP port or socket.")
	rootCmd.Flags().Bool("sync", false, "Make X calls synchronous.")
	rootCmd.Flags().BoolP("version", "v", false, "Print Pale Moon version.")

	rootCmd.Flag("display").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"display":          os.ActionDisplays(),
		"profile":          carapace.ActionFiles(),
		"recording":        carapace.ActionFiles(),
		"recording-output": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
