package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tor-browser",
	Short: "Tor Browser",
	Long:  "https://www.torproject.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("dir", "", "The Tor-Browser directory to use")
	rootCmd.Flags().BoolP("erase", "e", false, "Erase the copy in your home directory")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().BoolP("refresh", "r", false, "Refresh the copy in your home directory and launch tor-browser")
	rootCmd.Flags().BoolP("update", "u", false, "Search in AUR for a new release and install it")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
