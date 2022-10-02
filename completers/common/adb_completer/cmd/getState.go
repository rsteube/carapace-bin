package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getStateCmd = &cobra.Command{
	Use:   "get-state",
	Short: "print offline | bootloader | device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getStateCmd).Standalone()

	rootCmd.AddCommand(getStateCmd)
}
