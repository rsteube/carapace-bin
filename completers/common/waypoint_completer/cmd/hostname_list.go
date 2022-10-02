package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hostname_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered hostnames",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostname_listCmd).Standalone()

	addGlobalOptions(hostname_listCmd)

	hostnameCmd.AddCommand(hostname_listCmd)
}
