package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "Mirror a project/repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_mirrorCmd).Standalone()
	repo_mirrorCmd.Flags().Bool("allow-divergence", false, "Determines if divergent refs are skipped.")
	repo_mirrorCmd.Flags().String("direction", "pull", "Mirror direction")
	repo_mirrorCmd.Flags().Bool("enabled", true, "Determines if the mirror is enabled.")
	repo_mirrorCmd.Flags().Bool("protected-branches-only", false, "Determines if only protected branches are mirrored.")
	repo_mirrorCmd.Flags().String("url", "", "The target URL to which the repository is mirrored.")
	repoCmd.AddCommand(repo_mirrorCmd)

	carapace.Gen(repo_mirrorCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("pull", "push"),
	})

	carapace.Gen(repo_mirrorCmd).PositionalCompletion(
		action.ActionRepo(repo_mirrorCmd),
	)
}
