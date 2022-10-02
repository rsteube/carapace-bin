package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Displays a list of available conda commands and their help strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("help", "h", false, "Show this help message and exit.")
	helpCmd.Flags().Bool("help).", false, "")
	rootCmd.AddCommand(helpCmd)
}
