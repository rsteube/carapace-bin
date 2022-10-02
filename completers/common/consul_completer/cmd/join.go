package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Tell Consul agent to join cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(joinCmd).Standalone()
	addClientFlags(joinCmd)
	joinCmd.Flags().Bool("wan", false, "Joins a server to another server in the WAN pool.")

	rootCmd.AddCommand(joinCmd)
}
