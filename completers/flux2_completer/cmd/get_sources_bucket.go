package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sources_bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Get Bucket source statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sources_bucketCmd).Standalone()
	get_sourcesCmd.AddCommand(get_sources_bucketCmd)
}
