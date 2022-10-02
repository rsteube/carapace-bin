package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Bump a package version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	versionCmd.Flags().Bool("allow-same-version", false, "prevent error when same version as the current one is used")
	versionCmd.Flags().Bool("commit-hooks", false, "run git commit hooks")
	versionCmd.Flags().Bool("git-tag-version", false, "tag the commit")
	versionCmd.Flags().Bool("json", false, "output as json")
	versionCmd.Flags().String("prerelease-id", "", "prerelease identififer")
	versionCmd.Flags().Bool("sign-git-tag", false, "sign with gpg signature")
	addWorkspaceFlags(versionCmd)

	rootCmd.AddCommand(versionCmd)

	carapace.Gen(versionCmd).PositionalCompletion(
		carapace.ActionValues("major", "minor", "patch", "premajor", "preminor", "prepatch", "prerelease", "from-git"),
	)
}
