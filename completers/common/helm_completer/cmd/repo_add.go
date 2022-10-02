package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var repo_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a chart repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	repo_addCmd.Flags().Bool("allow-deprecated-repos", false, "by default, this command will not allow adding official repos that have been permanently deleted. This disables that behavior")
	repo_addCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	repo_addCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	repo_addCmd.Flags().Bool("force-update", false, "replace (overwrite) the repo if it already exists")
	repo_addCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the repository")
	repo_addCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	repo_addCmd.Flags().Bool("no-update", false, "Ignored. Formerly, it would disabled forced updates. It is deprecated by force-update.")
	repo_addCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	repo_addCmd.Flags().String("password", "", "chart repository password")
	repo_addCmd.Flags().String("username", "", "chart repository username")
	repoCmd.AddCommand(repo_addCmd)

	carapace.Gen(repo_addCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":   carapace.ActionFiles(),
		"cert-file": carapace.ActionFiles(),
		"key-file":  carapace.ActionFiles(),
	})
}
