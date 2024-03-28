package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debug_loadbalancerCmd = &cobra.Command{
	Use:     "loadbalancer",
	Short:   "Debug the loadbalancer",
	Aliases: []string{"lb"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_loadbalancerCmd).Standalone()

	debugCmd.AddCommand(debug_loadbalancerCmd)
}
