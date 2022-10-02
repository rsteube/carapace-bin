package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "google-chrome",
	Short: "chrome browser",
	Long:  "https://www.google.com/chrome/index.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteAlias(alias string) error {
	rootCmd.Use = alias
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("app", "", "Runs URL in app mode")
	rootCmd.Flags().Bool("new-window", false, "If PATH or URL is given, open it in a new window.")
	rootCmd.Flags().Bool("no-proxy-server", false, "Disables  the  proxy  server")
	rootCmd.Flags().String("password-store", "", "Set the password store to use")
	rootCmd.Flags().Bool("proxy-auto-detect", false, "Autodetect proxy configuration")
	rootCmd.Flags().String("proxy-pac-url", "", "Specify proxy autoconfiguration URL")
	rootCmd.Flags().String("proxy-server", "", "Specify the HTTP/SOCKS4/SOCKS5 proxy server to use for requests.")
	rootCmd.Flags().String("user-data-dir", "", "Specifies the directory that user data is  kept  in")
	rootCmd.Flags().Bool("version", false, "Show version information.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"password-store": carapace.ActionValues("basic", "gnome", "kwallet"),
		"user-data-dir":  carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
