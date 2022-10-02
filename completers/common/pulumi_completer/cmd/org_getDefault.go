package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var org_getDefaultCmd = &cobra.Command{
	Use:   "get-default",
	Short: "Get the default org for the current backend",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(org_getDefaultCmd).Standalone()
	orgCmd.AddCommand(org_getDefaultCmd)
}
