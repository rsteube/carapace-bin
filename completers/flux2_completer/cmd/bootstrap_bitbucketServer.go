package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_bitbucketServerCmd = &cobra.Command{
	Use:   "bitbucket-server",
	Short: "Bootstrap toolkit components in a Bitbucket Server repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_bitbucketServerCmd).Standalone()
	bootstrap_bitbucketServerCmd.Flags().StringSlice("group", []string{}, "Bitbucket Server groups to be given write access (also accepts comma-separated values)")
	bootstrap_bitbucketServerCmd.Flags().String("hostname", "", "Bitbucket Server hostname")
	bootstrap_bitbucketServerCmd.Flags().String("interval", "", "sync interval")
	bootstrap_bitbucketServerCmd.Flags().String("owner", "", "Bitbucket Server user or project name")
	bootstrap_bitbucketServerCmd.Flags().String("path", "", "path relative to the repository root, when specified the cluster sync will be scoped to this path")
	bootstrap_bitbucketServerCmd.Flags().Bool("personal", false, "if true, the owner is assumed to be a Bitbucket Server user; otherwise a group")
	bootstrap_bitbucketServerCmd.Flags().Bool("private", true, "if true, the repository is setup or configured as private")
	bootstrap_bitbucketServerCmd.Flags().Bool("read-write-key", false, "if true, the deploy key is configured with read/write permissions")
	bootstrap_bitbucketServerCmd.Flags().Bool("reconcile", false, "if true, the configured options are also reconciled if the repository already exists")
	bootstrap_bitbucketServerCmd.Flags().String("repository", "", "Bitbucket Server repository name")
	bootstrap_bitbucketServerCmd.Flags().StringP("username", "u", "git", "authentication username")
	bootstrapCmd.AddCommand(bootstrap_bitbucketServerCmd)
}
