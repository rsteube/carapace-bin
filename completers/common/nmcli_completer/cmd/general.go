package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generalCmd = &cobra.Command{
	Use:   "general",
	Short: "NetworkManager's general status and operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generalCmd).Standalone()

	rootCmd.AddCommand(generalCmd)
}
