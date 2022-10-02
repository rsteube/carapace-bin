package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-conflist",
	Short: "list the configuration files processed by Fontconfig",
	Long:  "https://manpages.debian.org/testing/fontconfig/fc-conflist.1.en.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("version", "V", false, "display font config version and exit")
}
