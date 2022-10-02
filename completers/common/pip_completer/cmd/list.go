package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("editable", "e", false, "List editable projects.")
	listCmd.Flags().String("exclude", "", "Exclude specified package from the output")
	listCmd.Flags().Bool("exclude-editable", false, "Exclude editable package from output.")
	listCmd.Flags().String("extra-index-url", "", "Extra URLs of package indexes to use in addition to --index-url.")
	listCmd.Flags().StringP("find-links", "f", "", "parse file for links to archives")
	listCmd.Flags().String("format", "", "Select the output format")
	listCmd.Flags().Bool("include-editable", false, "Include editable package from output.")
	listCmd.Flags().StringP("index-url", "i", "", "Base URL of the Python Package Index")
	listCmd.Flags().BoolP("local", "l", false, "do not list globally-installed packages.")
	listCmd.Flags().Bool("no-index", false, "Ignore package index")
	listCmd.Flags().Bool("not-required", false, "List packages that are not dependencies of installed packages.")
	listCmd.Flags().BoolP("outdated", "o", false, "List outdated packages")
	listCmd.Flags().String("path", "", "Restrict to the specified installation path for listing packages")
	listCmd.Flags().Bool("pre", false, "Include pre-release and development versions.")
	listCmd.Flags().BoolP("uptodate", "u", false, "List uptodate packages")
	listCmd.Flags().Bool("user", false, "Only output packages installed in user-site.")
	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"exclude":    pip.ActionInstalledPackages(),
		"find-links": carapace.ActionFiles(),
		"path":       carapace.ActionFiles(),
	})
}
