package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_helmreleaseCmd = &cobra.Command{
	Use:   "helmrelease",
	Short: "Export HelmRelease resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_helmreleaseCmd).Standalone()
	exportCmd.AddCommand(export_helmreleaseCmd)
}
