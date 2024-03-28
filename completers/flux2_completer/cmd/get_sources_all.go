package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sources_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all source statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sources_allCmd).Standalone()
	get_sourcesCmd.AddCommand(get_sources_allCmd)
}
