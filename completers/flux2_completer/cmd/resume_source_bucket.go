package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var resume_source_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Resume a suspended Bucket",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resume_source_bucketCmd).Standalone()
	resume_sourceCmd.AddCommand(resume_source_bucketCmd)
}
