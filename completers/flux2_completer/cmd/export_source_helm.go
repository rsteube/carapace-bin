package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Export HelmRepository sources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_source_helmCmd).Standalone()
	export_sourceCmd.AddCommand(export_source_helmCmd)
}
