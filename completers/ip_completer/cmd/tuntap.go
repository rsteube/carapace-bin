package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var tuntapCmd = &cobra.Command{
	Use:   "tuntap",
	Short: "manage TUN/TAP devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tuntapCmd).Standalone()

	rootCmd.AddCommand(tuntapCmd)
}
