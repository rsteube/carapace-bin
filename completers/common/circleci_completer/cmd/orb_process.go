package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_processCmd = &cobra.Command{
	Use:   "process",
	Short: "Validate an orb and print its form after all pre-registration processing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_processCmd).Standalone()
	orbCmd.AddCommand(orb_processCmd)
}
