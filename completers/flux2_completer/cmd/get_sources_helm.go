package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sources_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Get HelmRepository source statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sources_helmCmd).Standalone()
	get_sourcesCmd.AddCommand(get_sources_helmCmd)
}
