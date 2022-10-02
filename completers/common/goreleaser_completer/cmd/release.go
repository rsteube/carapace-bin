package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Releases the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.Flags().StringP("config", "f", "", "Load configuration from file")
	releaseCmd.Flags().Bool("debug", false, "Enable debug mode")
	releaseCmd.Flags().BoolP("help", "h", false, "help for release")
	releaseCmd.Flags().StringP("parallelism", "p", "", "Amount tasks to run concurrently (default 4)")
	releaseCmd.Flags().String("release-footer", "", "Load custom release notes footer from a markdown file")
	releaseCmd.Flags().String("release-header", "", "Load custom release notes header from a markdown file")
	releaseCmd.Flags().String("release-notes", "", "Load custom release notes from a markdown file")
	releaseCmd.Flags().Bool("rm-dist", false, "Remove the dist folder before building")
	releaseCmd.Flags().Bool("skip-publish", false, "Skips publishing artifacts")
	releaseCmd.Flags().Bool("skip-sign", false, "Skips signing the artifacts")
	releaseCmd.Flags().Bool("skip-validate", false, "Skips several sanity checks")
	releaseCmd.Flags().Bool("snapshot", false, "Generate an unversioned snapshot release, skipping all validations and without publishing any artifacts")
	releaseCmd.Flags().String("timeout", "", "Timeout to the entire release process (default 30m0s)")
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
