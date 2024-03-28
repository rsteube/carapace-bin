package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sourcesCmd = &cobra.Command{
	Use:   "sources",
	Short: "Get source statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sourcesCmd).Standalone()
	getCmd.AddCommand(get_sourcesCmd)
}
