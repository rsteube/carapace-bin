package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Request user instance or container exit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exitCmd).Standalone()

	rootCmd.AddCommand(exitCmd)
}
