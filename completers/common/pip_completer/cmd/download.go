package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downloadCmd).Standalone()

	downloadCmd.Flags().StringArray("abi", []string{}, "Only use wheels compatible with Python abi <abi>")
	downloadCmd.Flags().StringArrayP("constraint", "c", []string{}, "Constrain versions using the given constraints file")
	downloadCmd.Flags().StringP("dest", "d", "", "Download packages into <dir>")
	downloadCmd.Flags().String("extra-index-url", "", "Extra URLs of package indexes to use")
	downloadCmd.Flags().StringP("find-links", "f", "", "If a URL or path to an html file, then parse for links")
	downloadCmd.Flags().String("global-option", "", "Extra global options to be supplied to the setup.py call")
	downloadCmd.Flags().String("implementation", "", "Only use wheels compatible with Python implementation <implementation>")
	downloadCmd.Flags().StringP("index-url", "i", "", "Base URL of the Python Package Index")
	downloadCmd.Flags().StringArray("no-binary", []string{}, "Do not use binary packages")
	downloadCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building a modern source distribution")
	downloadCmd.Flags().Bool("no-clean", false, "Don't clean up build directories.")
	downloadCmd.Flags().Bool("no-deps", false, "Don't install package dependencies")
	downloadCmd.Flags().Bool("no-index", false, "Ignore package index")
	downloadCmd.Flags().StringArray("only-binary", []string{}, "Do not use source packages")
	downloadCmd.Flags().StringArray("platform", []string{}, "Only use wheels compatible with <platform>")
	downloadCmd.Flags().Bool("pre", false, "Include pre-release and development versions")
	downloadCmd.Flags().Bool("prefer-binary", false, "Prefer older binary packages over newer source packages")
	downloadCmd.Flags().String("progress-bar", "", "Specify type of progress to be displayed")
	downloadCmd.Flags().String("python-version", "", "The Python interpreter version to use for wheel")
	downloadCmd.Flags().Bool("require-hashes", false, "Require a hash to check each requirement against")
	downloadCmd.Flags().StringArrayP("requirement", "r", []string{}, "Install from the given requirements file")
	downloadCmd.Flags().String("src", "", "Directory to check out editable projects into")
	downloadCmd.Flags().Bool("use-pep517", false, "Use PEP 517 for building source distributions")
	rootCmd.AddCommand(downloadCmd)

	carapace.Gen(downloadCmd).FlagCompletion(carapace.ActionMap{
		"constraint":     carapace.ActionFiles(),
		"find-links":     carapace.ActionFiles(),
		"implementation": carapace.ActionValues("pp", "jy", "cp", "ip"),
		"progress-bar":   carapace.ActionValues("off", "on", "ascii", "pretty", "emoji"),
		"requirement":    carapace.ActionFiles(),
		"src":            carapace.ActionDirectories(),
	})

	carapace.Gen(downloadCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO support multiple index urls (--index-url) and update caching accordingly
			return pip.ActionPackageSearch()
		}),
	)

}
