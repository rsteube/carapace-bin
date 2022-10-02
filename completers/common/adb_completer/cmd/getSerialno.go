package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var getSerialnoCmd = &cobra.Command{
	Use:   "get-serialno",
	Short: "print <serial-number>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getSerialnoCmd).Standalone()

	rootCmd.AddCommand(getSerialnoCmd)
}
