package cmd

import (
	"github.com/spf13/cobra"
)

var mod_npmCmd = &cobra.Command{
	Use:   "npm",
	Short: "Various npm helpers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.AddCommand(mod_npmCmd)
}
