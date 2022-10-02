package cmd

import (
	"github.com/spf13/cobra"
)

var mod_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Resolves dependencies in your current Hugo Project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.AddCommand(mod_getCmd)
}
