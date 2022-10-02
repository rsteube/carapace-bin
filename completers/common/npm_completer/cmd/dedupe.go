package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var dedupeCmd = &cobra.Command{
	Use:   "dedupe",
	Short: "Reduce duplication in the package tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dedupeCmd).Standalone()

	dedupeCmd.Flags().Bool("audit", false, "submit audit reports")
	dedupeCmd.Flags().Bool("bin-links", false, "create symlinks")
	dedupeCmd.Flags().Bool("dry-run", false, "only report changes")
	dedupeCmd.Flags().Bool("fund", false, "show funding message")
	dedupeCmd.Flags().Bool("global-style", false, "use global layout")
	dedupeCmd.Flags().Bool("ignore-scripts", false, "do not run scripts specified in package.json")
	dedupeCmd.Flags().Bool("legacy-bundling", false, "support older npm version")
	dedupeCmd.Flags().String("omit", "", "omit dependency types")
	dedupeCmd.Flags().Bool("package-lock", false, "when false ignore `package-lock.json`")
	dedupeCmd.Flags().Bool("strict-peer-deps", false, "any conflicting `peerDependencies` will be treated as install failure")
	addWorkspaceFlags(dedupeCmd)
	rootCmd.AddCommand(dedupeCmd)
}
