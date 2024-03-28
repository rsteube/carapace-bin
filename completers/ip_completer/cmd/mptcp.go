package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mptcpCmd = &cobra.Command{
	Use:   "mptcp",
	Short: "manage MPTCP path manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mptcpCmd).Standalone()

	rootCmd.AddCommand(mptcpCmd)
}
