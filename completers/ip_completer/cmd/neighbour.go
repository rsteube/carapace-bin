package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var neighbourCmd = &cobra.Command{
	Use:   "neighbour",
	Short: "manage ARP or NDISC cache entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neighbourCmd).Standalone()

	rootCmd.AddCommand(neighbourCmd)
}
