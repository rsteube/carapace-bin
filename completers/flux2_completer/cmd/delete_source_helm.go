package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_source_helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Delete a HelmRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_source_helmCmd).Standalone()
	delete_sourceCmd.AddCommand(delete_source_helmCmd)
}
