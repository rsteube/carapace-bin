package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspend_source_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Suspend reconciliation of a Bucket",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspend_source_bucketCmd).Standalone()
	suspend_sourceCmd.AddCommand(suspend_source_bucketCmd)
}
