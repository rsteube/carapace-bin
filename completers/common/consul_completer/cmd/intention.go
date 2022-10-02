package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intentionCmd = &cobra.Command{
	Use:   "intention",
	Short: "Interact with Connect service intentions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intentionCmd).Standalone()

	rootCmd.AddCommand(intentionCmd)
}
