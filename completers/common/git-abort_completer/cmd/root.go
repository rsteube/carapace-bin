package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-abort",
	Short: "Abort current rebase, merge or cherry-pick, without the need to find exact command in history",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-abort",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "show help")
}
