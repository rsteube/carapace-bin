package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var leaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Gracefully leaves the Consul cluster and shuts down",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(leaveCmd).Standalone()
	addClientFlags(leaveCmd)

	rootCmd.AddCommand(leaveCmd)
}
