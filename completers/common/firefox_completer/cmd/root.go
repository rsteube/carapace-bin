package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "firefox",
	Short: "Firefox Browser",
	Long:  "https://www.mozilla.org/en-US/firefox/new/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("P", "P", "", "Start with <profile>.")
	rootCmd.Flags().Bool("ProfileManager", false, "Start with ProfileManager.")
	rootCmd.Flags().Bool("allow-downgrade", false, "Allows downgrading a profile.")
	rootCmd.Flags().Bool("browser", false, "Open a browser window.")
	rootCmd.Flags().Bool("devtools", false, "Open DevTools on initial load.")
	rootCmd.Flags().String("display", "", "X display to use")
	rootCmd.Flags().Bool("first-startup", false, "Run post-install actions before opening a new window.")
	rootCmd.Flags().Bool("full-version", false, "Print Firefox version, build and platform build ids.")
	rootCmd.Flags().Bool("headless", false, "Run without a GUI.")
	rootCmd.Flags().BoolP("help", "h", false, "Print this message.")
	rootCmd.Flags().Bool("jsconsole", false, "Open the Browser Console.")
	rootCmd.Flags().Bool("migration", false, "Start with migration wizard.")
	rootCmd.Flags().Bool("new-instance", false, "Open new instance, not a new window in running instance.")
	rootCmd.Flags().String("new-tab", "", "Open <url> in a new tab.")
	rootCmd.Flags().Bool("no-remote", false, "Do not accept or send remote commands")
	rootCmd.Flags().Bool("preferences", false, "Open Preferences dialog.")
	rootCmd.Flags().String("profile", "", "Start with profile at <path>.")
	rootCmd.Flags().Bool("safe-mode", false, "Disables extensions and themes for this session.")
	rootCmd.Flags().String("search", "", "Search <term> with your default search engine.")
	rootCmd.Flags().String("ssb", "", "Open a site specific browser for <uri>.")
	rootCmd.Flags().String("start-debugger-server", "", "Start the devtools server")
	rootCmd.Flags().Bool("sync", false, "Make X calls synchronous")
	rootCmd.Flags().BoolP("version", "v", false, "Print Firefox version.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
