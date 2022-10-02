package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "List jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listJobsCmd).Standalone()

	rootCmd.AddCommand(listJobsCmd)
}
