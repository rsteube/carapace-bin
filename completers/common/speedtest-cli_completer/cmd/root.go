package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/speedtest-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "speedtest-cli",
	Short: "Command line interface for testing internet bandwidth using speedtest.net",
	Long:  "https://www.speedtest.net/apps/cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bytes", false, "Display values in bytes instead of bits.")
	rootCmd.Flags().Bool("csv", false, "Suppress verbose output, only show basic information in CSV format.")
	rootCmd.Flags().String("csv-delimiter", "", "Single character delimiter to use in CSV output.")
	rootCmd.Flags().Bool("csv-header", false, "Print CSV headers")
	rootCmd.Flags().StringArray("exclude", []string{}, "Exclude a server from selection.")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().Bool("json", false, "Suppress verbose output, only show basic information in JSON format.")
	rootCmd.Flags().Bool("list", false, "Display a list of speedtest.net servers sorted by distance")
	rootCmd.Flags().String("mini", "", "URL of the Speedtest Mini server")
	rootCmd.Flags().Bool("no-download", false, "Do not perform download test")
	rootCmd.Flags().Bool("no-pre-allocate", false, "Do not pre allocate upload data.")
	rootCmd.Flags().Bool("no-upload", false, "Do not perform upload test")
	rootCmd.Flags().Bool("secure", false, "Use HTTPS instead of HTTP when communicating with speedtest.net operated servers")
	rootCmd.Flags().StringArray("server", []string{}, "Specify a server ID to test against.")
	rootCmd.Flags().Bool("share", false, "Generate and provide a URL to the speedtest.net share results image")
	rootCmd.Flags().Bool("simple", false, "Suppress verbose output, only show basic information")
	rootCmd.Flags().Bool("single", false, "Only use a single connection instead of multiple.")
	rootCmd.Flags().String("source", "", "Source IP address to bind to")
	rootCmd.Flags().String("timeout", "", "HTTP timeout in seconds. Default 10")
	rootCmd.Flags().Bool("version", false, "Show the version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exclude": action.ActionServers(),
		"server":  action.ActionServers(),
	})
}
