package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cargo-set-version",
	Short: "Change a package's version in the local manifest file",
	Long:  "https://github.com/killercup/cargo-edit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "[deprecated in favor of `--workspace`]")
	rootCmd.Flags().String("bump", "", "Increment manifest version")
	rootCmd.Flags().Bool("dry-run", false, "Print changes to be made without making them")
	rootCmd.Flags().String("exclude", "", "Crates to exclude and not modify")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().String("manifest-path", "", "Path to the manifest to upgrade")
	rootCmd.Flags().StringP("metadata", "m", "", "Specify the version metadata field (e.g. a wrapped libraries version)")
	rootCmd.Flags().StringP("package", "p", "", "Package id of the crate to change the version of")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.Flags().Bool("workspace", false, "Modify all packages in the workspace")

	// TODO flag completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bump":          carapace.ActionValues("major", "minor", "patch", "release", "rc", "beta", "alpha"),
		"manifest-path": carapace.ActionFiles(".toml"),
	})

	// TODO complete current version as positional for easy edit
}
