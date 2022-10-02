package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Manage volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volumeCmd).Standalone()

	rootCmd.AddCommand(volumeCmd)
}
