package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pubCmd = &cobra.Command{
	Use:   "pub",
	Short: "Work with packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pubCmd).Standalone()

	pubCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pubCmd.Flags().Bool("no-trace", false, "Do not print debugging information when an error occurs.")
	pubCmd.Flags().Bool("trace", false, "Print debugging information when an error occurs.")
	pubCmd.Flags().BoolP("verbose", "v", false, "Shortcut for \"--verbosity=all\".")
	rootCmd.AddCommand(pubCmd)
}
