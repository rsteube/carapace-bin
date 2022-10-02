package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize new module in current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_initCmd).Standalone()

	modCmd.AddCommand(mod_initCmd)
}
