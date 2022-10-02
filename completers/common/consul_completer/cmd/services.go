package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Interact with services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicesCmd).Standalone()

	rootCmd.AddCommand(servicesCmd)
}
