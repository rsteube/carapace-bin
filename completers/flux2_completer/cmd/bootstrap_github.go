package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Bootstrap toolkit components in a GitHub repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_githubCmd).Standalone()
	bootstrap_githubCmd.Flags().String("hostname", "github.com", "GitHub hostname")
	bootstrap_githubCmd.Flags().String("interval", "", "sync interval")
	bootstrap_githubCmd.Flags().String("owner", "", "GitHub user or organization name")
	bootstrap_githubCmd.Flags().String("path", "", "path relative to the repository root, when specified the cluster sync will be scoped to this path")
	bootstrap_githubCmd.Flags().Bool("personal", false, "if true, the owner is assumed to be a GitHub user; otherwise an org")
	bootstrap_githubCmd.Flags().Bool("private", true, "if true, the repository is setup or configured as private")
	bootstrap_githubCmd.Flags().Bool("read-write-key", false, "if true, the deploy key is configured with read/write permissions")
	bootstrap_githubCmd.Flags().Bool("reconcile", false, "if true, the configured options are also reconciled if the repository already exists")
	bootstrap_githubCmd.Flags().String("repository", "", "GitHub repository name")
	bootstrap_githubCmd.Flags().StringSlice("team", []string{}, "GitHub team and the access to be given to it(team:maintain). Defaults to maintainer access if no access level is specified (also accepts comma-separated values)")
	bootstrapCmd.AddCommand(bootstrap_githubCmd)
}
