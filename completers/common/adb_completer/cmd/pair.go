package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pairCmd = &cobra.Command{
	Use:   "pair",
	Short: "pair with a device for secure TCP/IP communication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pairCmd).Standalone()

	rootCmd.AddCommand(pairCmd)
}
