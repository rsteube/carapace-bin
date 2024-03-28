package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_sources_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Get GitRepository source statuses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_sources_gitCmd).Standalone()
	get_sourcesCmd.AddCommand(get_sources_gitCmd)
}
