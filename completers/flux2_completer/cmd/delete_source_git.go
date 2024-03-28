package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_source_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Delete a GitRepository source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_source_gitCmd).Standalone()
	delete_sourceCmd.AddCommand(delete_source_gitCmd)
}
