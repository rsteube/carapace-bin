package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade deno executable to the given version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("canary", false, "Upgrade to canary builds")
	upgradeCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	upgradeCmd.Flags().Bool("dry-run", false, "Perform all checks without replacing old exe")
	upgradeCmd.Flags().BoolP("force", "f", false, "Replace current exe even if not out-of-date")
	upgradeCmd.Flags().BoolP("help", "h", false, "Print help information")
	upgradeCmd.Flags().String("output", "", "The path to output the updated version to")
	upgradeCmd.Flags().BoolP("quiet", "q", false, "Suppress diagnostic output")
	upgradeCmd.Flags().Bool("unstable", false, "Enable unstable features and APIs")
	upgradeCmd.Flags().String("version", "", "The version to upgrade to")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"cert":   carapace.ActionFiles(),
		"output": carapace.ActionDirectories(),
	})
}
