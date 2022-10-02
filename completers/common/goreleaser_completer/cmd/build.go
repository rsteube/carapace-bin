package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().StringP("config", "f", "", "Load configuration from file")
	buildCmd.Flags().Bool("debug", false, "Enable debug mode")
	buildCmd.Flags().BoolP("help", "h", false, "help for build")
	buildCmd.Flags().StringP("parallelism", "p", "", "Amount tasks to run concurrently (default 4)")
	buildCmd.Flags().Bool("rm-dist", false, "Remove the dist folder before building")
	buildCmd.Flags().Bool("skip-post-hooks", false, "Skips all post-build hooks")
	buildCmd.Flags().Bool("skip-validate", false, "Skips several sanity checks")
	buildCmd.Flags().Bool("snapshot", false, "Generate an unversioned snapshot build, skipping all validations and without publishing any artifacts")
	buildCmd.Flags().String("timeout", "", "Timeout to the entire build process (default 30m0s)")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
