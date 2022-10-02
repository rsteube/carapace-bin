package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "list all command names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandsCmd).Standalone()

	rootCmd.AddCommand(commandsCmd)
}
