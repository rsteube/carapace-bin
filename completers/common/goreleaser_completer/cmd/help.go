package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Help about any command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().Bool("debug", false, "Enable debug mode")
	helpCmd.Flags().BoolP("help", "h", false, "help for help")
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.ActionValues("build", "check", "completion", "init", "release"),
	)
}
