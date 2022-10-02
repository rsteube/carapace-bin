package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize declared variables in waypoint.hcl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_syncCmd).Standalone()

	addGlobalOptions(config_syncCmd)

	configCmd.AddCommand(config_syncCmd)
}
