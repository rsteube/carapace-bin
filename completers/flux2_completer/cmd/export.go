package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export resources in YAML format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()
	exportCmd.PersistentFlags().Bool("all", false, "select all resources")
	rootCmd.AddCommand(exportCmd)
}
