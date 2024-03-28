package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var addrlabelCmd = &cobra.Command{
	Use:   "addrlabel",
	Short: "label configuration for protocol address selection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addrlabelCmd).Standalone()

	rootCmd.AddCommand(addrlabelCmd)
}
