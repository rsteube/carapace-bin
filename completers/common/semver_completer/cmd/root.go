package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "A JavaScript implementation of the https://semver.org/ specification",
	Long:  "https://github.com/npm/node-semver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("coerce", "c", false, "Coerce a string into SemVer if possible")
	rootCmd.Flags().BoolP("include-prerelease", "p", false, "Always include prerelease versions in range matching")
	rootCmd.Flags().StringP("increment", "i", "", "Increment a version by the specified level.")
	rootCmd.Flags().BoolP("loose", "l", false, "Interpret versions and ranges loosely")
	rootCmd.Flags().Bool("ltr", false, "Coerce version strings left to right (default)")
	rootCmd.Flags().String("preid", "", "Identifier to be used to prefix version increments.")
	rootCmd.Flags().StringP("range", "r", "", "Print versions that match the specified range.")
	rootCmd.Flags().Bool("rtl", false, "Coerce version strings right to left")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"increment": carapace.ActionValues("major", "minor", "patch", "premajor", "preminor", "prepatch", "prerelease"),
		"preid":     carapace.ActionValues("premajor", "preminor", "prepatch", "prerelease"),
	})
}
