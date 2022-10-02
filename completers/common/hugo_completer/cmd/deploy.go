package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hugo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your site to a Cloud provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	deployCmd.Flags().Bool("confirm", false, "ask for confirmation before making changes to the target")
	deployCmd.Flags().Bool("dryRun", false, "dry run")
	deployCmd.Flags().Bool("force", false, "force upload of all files")
	deployCmd.Flags().Bool("invalidateCDN", true, "invalidate the CDN cache listed in the deployment target")
	deployCmd.Flags().Int("maxDeletes", 256, "maximum # of files to delete, or -1 to disable")
	deployCmd.Flags().String("target", "", "target deployment from deployments section in config file; defaults to the first one")
	rootCmd.AddCommand(deployCmd)

	carapace.Gen(deployCmd).FlagCompletion(carapace.ActionMap{
		"target": action.ActionDeployments(),
	})
}
