package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gftp",
	Short: "file transfer client for *NIX based machines",
	Long:  "https://github.com/masneyb/gftp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display program usage.")
	rootCmd.Flags().Bool("info", false, "Display some information about how gFTP was built.")
	rootCmd.Flags().BoolP("version", "v", false, "Display the current version of gFTP.")
}
