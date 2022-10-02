package cmd

import (
	"github.com/spf13/cobra"
)

var mod_vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Vendor all module dependencies into the _vendor directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.AddCommand(mod_vendorCmd)
}
