package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_receiversCmd = &cobra.Command{
	Use:   "receivers",
	Short: "Get Receiver statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_receiversCmd).Standalone()
	getCmd.AddCommand(get_receiversCmd)
}
