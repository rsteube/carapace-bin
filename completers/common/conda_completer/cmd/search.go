package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for packages and display associated information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().Bool("envs", false, "Search all of the current user's environments.")
	searchCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	searchCmd.Flags().BoolP("info", "i", false, "Provide detailed information about each package.")
	searchCmd.Flags().BoolP("insecure", "k", false, "Allow conda to perform \"insecure\" SSL connections and transfers.")
	searchCmd.Flags().Bool("json", false, "Report all output as json.")
	searchCmd.Flags().Bool("offline", false, "Offline mode.")
	searchCmd.Flags().Bool("override-channels", false, "Do not search default or .condarc channels.")
	searchCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar.")
	searchCmd.Flags().String("repodata-fn", "", "Specify name of repodata on remote server.")
	searchCmd.Flags().BoolP("use-index-cache", "C", false, "Use cache of channel index files, even if it has expired.")
	searchCmd.Flags().Bool("use-local", false, "Use locally built packages.")
	searchCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace.")
	rootCmd.AddCommand(searchCmd)
}
