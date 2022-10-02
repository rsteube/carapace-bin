package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sleep",
	Short: "delay for a specified amount of time",
	Long:  "https://linux.die.net/man/3/sleep",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
