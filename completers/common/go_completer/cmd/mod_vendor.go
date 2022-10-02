package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "make vendored copy of dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_vendorCmd).Standalone()

	mod_vendorCmd.Flags().BoolS("v", "v", false, "print the names of vendored modules and packages")
	modCmd.AddCommand(mod_vendorCmd)
}
