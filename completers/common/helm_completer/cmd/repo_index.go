package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_indexCmd = &cobra.Command{
	Use:   "index",
	Short: "generate an index file given a directory containing packaged charts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repo_indexCmd.Flags().String("merge", "", "merge the generated index into the given index")
	repo_indexCmd.Flags().String("url", "", "url of chart repository")
	repoCmd.AddCommand(repo_indexCmd)

	carapace.Gen(repo_indexCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
