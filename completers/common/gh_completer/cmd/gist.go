package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist",
	Short: "Manage gists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gistCmd).Standalone()
	rootCmd.AddCommand(gistCmd)
}
