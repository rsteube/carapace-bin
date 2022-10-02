package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Prepare a subset of your monorepo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneCmd).Standalone()

	pruneCmd.Flags().Bool("docker", false, "Output pruned workspace into 'full' and 'json' directories optimized for Docker layer caching.")
	pruneCmd.Flags().String("out-dir", "", "Set the root directory for files output by this command (default \"out\")")
	pruneCmd.Flags().String("scope", "", "Specify package to act as entry point for pruned monorepo (required).")
	rootCmd.AddCommand(pruneCmd)

	pruneCmd.Flag("out-dir").NoOptDefVal = " "
	pruneCmd.Flag("scope").NoOptDefVal = " "

	// TODO scope
	carapace.Gen(pruneCmd).FlagCompletion(carapace.ActionMap{
		"out-dir": carapace.ActionDirectories(),
	})
}
