package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sources_chartCmd = &cobra.Command{
	Use:   "chart",
	Short: "Get HelmChart statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sources_chartCmd).Standalone()
	get_sourcesCmd.AddCommand(get_sources_chartCmd)
}
