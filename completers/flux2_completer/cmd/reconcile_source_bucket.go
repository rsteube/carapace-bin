package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reconcile_source_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Reconcile a Bucket source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reconcile_source_bucketCmd).Standalone()
	reconcile_sourceCmd.AddCommand(reconcile_source_bucketCmd)
}
