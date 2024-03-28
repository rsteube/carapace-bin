package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Export Alert resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_alertCmd).Standalone()
	exportCmd.AddCommand(export_alertCmd)
}
