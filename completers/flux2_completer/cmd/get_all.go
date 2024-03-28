package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all resources and statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_allCmd).Standalone()
	getCmd.AddCommand(get_allCmd)
}
