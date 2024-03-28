package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Bootstrap toolkit components in a GitLab repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_gitlabCmd).Standalone()
	bootstrap_gitlabCmd.Flags().String("hostname", "gitlab.com", "GitLab hostname")
	bootstrap_gitlabCmd.Flags().String("interval", "", "sync interval")
	bootstrap_gitlabCmd.Flags().String("owner", "", "GitLab user or group name")
	bootstrap_gitlabCmd.Flags().String("path", "", "path relative to the repository root, when specified the cluster sync will be scoped to this path")
	bootstrap_gitlabCmd.Flags().Bool("personal", false, "if true, the owner is assumed to be a GitLab user; otherwise a group")
	bootstrap_gitlabCmd.Flags().Bool("private", true, "if true, the repository is setup or configured as private")
	bootstrap_gitlabCmd.Flags().Bool("read-write-key", false, "if true, the deploy key is configured with read/write permissions")
	bootstrap_gitlabCmd.Flags().Bool("reconcile", false, "if true, the configured options are also reconciled if the repository already exists")
	bootstrap_gitlabCmd.Flags().String("repository", "", "GitLab repository name")
	bootstrap_gitlabCmd.Flags().StringSlice("team", []string{}, "GitLab teams to be given maintainer access (also accepts comma-separated values)")
	bootstrapCmd.AddCommand(bootstrap_gitlabCmd)
}
