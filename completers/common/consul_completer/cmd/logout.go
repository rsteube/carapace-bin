package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Gracefully leaves the Consul cluster and shuts down",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()
	addClientFlags(logoutCmd)
	addServerFlags(logoutCmd)

	rootCmd.AddCommand(logoutCmd)
}
