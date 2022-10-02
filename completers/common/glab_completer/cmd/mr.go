package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mrCmd = &cobra.Command{
	Use:   "mr",
	Short: "Create, view and manage merge requests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mrCmd).Standalone()
	mrCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(mrCmd)

	carapace.Gen(mrCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(mrCmd),
	})
}
