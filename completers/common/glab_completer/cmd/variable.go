package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var variableCmd = &cobra.Command{
	Use:   "variable",
	Short: "Manage GitLab Project and Group Variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(variableCmd).Standalone()
	variableCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format or full URL or git URL")
	rootCmd.AddCommand(variableCmd)

	carapace.Gen(variableCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(variableCmd),
	})
}
