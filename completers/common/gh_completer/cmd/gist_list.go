package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gist_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your gists",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_listCmd).Standalone()
	gist_listCmd.Flags().IntP("limit", "L", 10, "Maximum number of gists to fetch")
	gist_listCmd.Flags().Bool("public", false, "Show only public gists")
	gist_listCmd.Flags().Bool("secret", false, "Show only secret gists")
	gistCmd.AddCommand(gist_listCmd)
}
