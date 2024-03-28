package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Export sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_sourceCmd).Standalone()
	export_sourceCmd.PersistentFlags().Bool("with-credentials", false, "include credential secrets")
	exportCmd.AddCommand(export_sourceCmd)
}
