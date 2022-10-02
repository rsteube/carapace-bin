package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var haltCmd = &cobra.Command{
	Use:   "halt",
	Short: "Shut down and halt the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(haltCmd).Standalone()

	rootCmd.AddCommand(haltCmd)
}
