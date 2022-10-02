package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mod_tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "add missing and remove unused modules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_tidyCmd).Standalone()

	mod_tidyCmd.Flags().BoolS("v", "v", false, "print information about removed modules")
	modCmd.AddCommand(mod_tidyCmd)
}
