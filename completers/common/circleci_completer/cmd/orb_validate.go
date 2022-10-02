package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orb_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an orb.yml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_validateCmd).Standalone()
	orbCmd.AddCommand(orb_validateCmd)
}
