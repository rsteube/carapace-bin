package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_alertsCmd = &cobra.Command{
	Use:   "alerts",
	Short: "Get Alert statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_alertsCmd).Standalone()
	getCmd.AddCommand(get_alertsCmd)
}
