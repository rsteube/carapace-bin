package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var org_setDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set the default organization for the current backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_setDefaultCmd).Standalone()
	orgCmd.AddCommand(org_setDefaultCmd)
}
