package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var maddressCmd = &cobra.Command{
	Use:   "maddress",
	Short: "multicast address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maddressCmd).Standalone()

	rootCmd.AddCommand(maddressCmd)
}
