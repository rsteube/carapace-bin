package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Bootstrap toolkit components in a Git repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_gitCmd).Standalone()
	bootstrap_gitCmd.Flags().String("interval", "", "sync interval")
	bootstrap_gitCmd.Flags().StringP("password", "p", "", "basic authentication password")
	bootstrap_gitCmd.Flags().String("path", "", "path relative to the repository root, when specified the cluster sync will be scoped to this path")
	bootstrap_gitCmd.Flags().BoolP("silent", "s", false, "assumes the deploy key is already setup, skips confirmation")
	bootstrap_gitCmd.Flags().String("url", "", "Git repository URL")
	bootstrap_gitCmd.Flags().StringP("username", "u", "git", "basic authentication username")
	bootstrapCmd.AddCommand(bootstrap_gitCmd)
}
