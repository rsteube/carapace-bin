package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_source_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Delete a Bucket source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_source_bucketCmd).Standalone()
	delete_sourceCmd.AddCommand(delete_source_bucketCmd)
}
