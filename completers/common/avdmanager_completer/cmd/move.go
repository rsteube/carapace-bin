package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Moves or renames an Android Virtual Device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	rootCmd.AddCommand(moveCmd)
}
