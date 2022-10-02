package cmd

import (
	"github.com/spf13/cobra"
)

var mod_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize this project as a Hugo Module.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.AddCommand(mod_initCmd)
}
