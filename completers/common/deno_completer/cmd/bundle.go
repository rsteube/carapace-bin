package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Bundle module and dependencies into single file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bundleCmd).Standalone()

	bundleCmd.Flags().String("cert", "", "Load certificate authority from PEM encoded file")
	bundleCmd.Flags().StringP("config", "c", "", "Load configuration file")
	bundleCmd.Flags().String("import-map", "", "Load import map file")
	bundleCmd.Flags().String("lock", "", "Check the specified lock file")
	bundleCmd.Flags().Bool("lock-write", false, "Write lock file (use with --lock)")
	bundleCmd.Flags().Bool("no-check", false, "Skip type checking modules")
	bundleCmd.Flags().Bool("no-remote", false, "Do not resolve remote modules")
	bundleCmd.Flags().StringP("reload", "r", "", "Reload source code cache (recompile TypeScript)")
	bundleCmd.Flags().Bool("watch", false, "UNSTABLE: Watch for file changes and restart process automatically")
	rootCmd.AddCommand(bundleCmd)

	bundleCmd.Flag("reload").NoOptDefVal = " "

	carapace.Gen(bundleCmd).FlagCompletion(carapace.ActionMap{
		"cert":       carapace.ActionFiles(),
		"config":     carapace.ActionFiles(),
		"import-map": carapace.ActionFiles(),
		"lock":       carapace.ActionFiles(),
	})

	carapace.Gen(bundleCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
