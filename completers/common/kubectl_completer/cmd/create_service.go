package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Create a service using specified subcommand",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_serviceCmd).Standalone()

	createCmd.AddCommand(create_serviceCmd)
}
