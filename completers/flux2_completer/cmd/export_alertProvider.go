package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_alertProviderCmd = &cobra.Command{
	Use:   "alert-provider",
	Short: "Export Provider resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_alertProviderCmd).Standalone()
	exportCmd.AddCommand(export_alertProviderCmd)
}
