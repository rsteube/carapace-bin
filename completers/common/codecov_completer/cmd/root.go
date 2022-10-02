package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codecov",
	Short: "codecov uploader",
	Long:  "https://docs.codecov.com/docs/codecov-uploader",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("CL", false, "Display a link for the current changelog")
	rootCmd.Flags().StringP("branch", "B", "", "Specify the branch manually")
	rootCmd.Flags().StringP("build", "b", "", "Specify the build number manually")
	rootCmd.Flags().Bool("changelog", false, "Display a link for the current changelog")
	rootCmd.Flags().BoolP("clean", "c", false, "Move discovered coverage reports to the trash")
	rootCmd.Flags().StringP("dir", "s", "", "Directory to search for coverage reports.")
	rootCmd.Flags().BoolP("dryRun", "d", false, "Don't upload files to Codecov")
	rootCmd.Flags().StringP("env", "e", "", "Specify environment variables to be included with this build.")
	rootCmd.Flags().StringP("feature", "X", "", "Toggle functionalities")
	rootCmd.Flags().StringP("file", "f", "", "Target file(s) to upload")
	rootCmd.Flags().BoolP("flags", "F", false, "Flag the upload to group coverage metrics")
	rootCmd.Flags().BoolP("help", "h", false, "Show help")
	rootCmd.Flags().StringP("name", "n", "", "Custom defined name of the upload.")
	rootCmd.Flags().BoolP("nonZero", "Z", false, "Should errors exit with a non-zero")
	rootCmd.Flags().StringP("parent", "N", "", "The commit SHA of the parent for which you are uploading coverage.")
	rootCmd.Flags().StringP("pr", "P", "", "Specify the pull request number mannually")
	rootCmd.Flags().StringP("rootDir", "R", "", "Specify the project root directory when not in a git repo")
	rootCmd.Flags().StringP("sha", "C", "", "Specify the commit SHA mannually")
	rootCmd.Flags().StringP("slug", "r", "", "Specify the slug manually (Enterprise use)")
	rootCmd.Flags().StringP("source", "Q", "", "Used internally by Codecov, this argument helps track wrappers of the uploader")
	rootCmd.Flags().StringP("tag", "T", "", "Specify the git tag")
	rootCmd.Flags().StringP("token", "t", "", "Codecov upload token")
	rootCmd.Flags().StringP("url", "u", "", "Change the upload host (Enterprise use)")
	rootCmd.Flags().BoolP("verbose", "v", false, "Run with verbose logging")
	rootCmd.Flags().Bool("version", false, "Show version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO rootDir for git actions
		"branch":  git.ActionRefs(git.RefOption{LocalBranches: true}),
		"dir":     carapace.ActionDirectories(),
		"file":    carapace.ActionFiles(),
		"rootDir": carapace.ActionDirectories(),
		"tag":     git.ActionRefs(git.RefOption{Tags: true}),
	})
}
