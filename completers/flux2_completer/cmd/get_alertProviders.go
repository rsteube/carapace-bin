package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_alertProvidersCmd = &cobra.Command{
	Use:   "alert-providers",
	Short: "Get Provider statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_alertProvidersCmd).Standalone()
	getCmd.AddCommand(get_alertProvidersCmd)
}
