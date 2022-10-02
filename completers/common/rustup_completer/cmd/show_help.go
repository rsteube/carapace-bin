package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var show_helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_helpCmd).Standalone()

	showCmd.AddCommand(show_helpCmd)
}
