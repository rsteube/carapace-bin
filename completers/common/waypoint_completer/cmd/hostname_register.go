package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var hostname_registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a hostname to route to your apps",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostname_registerCmd).Standalone()

	addGlobalOptions(hostname_registerCmd)

	hostnameCmd.AddCommand(hostname_registerCmd)
}
