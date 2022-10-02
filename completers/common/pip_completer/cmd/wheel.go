package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var wheelCmd = &cobra.Command{
	Use:   "wheel",
	Short: "Build Wheel archives for your requirements and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wheelCmd).Standalone()

	wheelCmd.Flags().String("build-option", "", "Extra arguments to be supplied to 'setup.py bdist_wheel'.")
	wheelCmd.Flags().StringP("constraint", "c", "", "Constrain versions using the given constraints file.")
	wheelCmd.Flags().StringP("editable", "e", "", "Install a project in editable mode from path.")
	wheelCmd.Flags().String("extra-index-url", "", "Extra URLs of package indexes to use in addition to --index-url.")
	wheelCmd.Flags().StringP("find-links", "f", "", "parse file for links to archives")
	wheelCmd.Flags().String("global-option", "", "Extra global options to be supplied to the setup.py call.")
	wheelCmd.Flags().Bool("ignore-requires-python", false, "Ignore the Requires-Python information.")
	wheelCmd.Flags().StringP("index-url", "i", "", "Base URL of the Python Package Index")
	wheelCmd.Flags().String("no-binary", "", "Do not use binary packages.")
	wheelCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building a modern source distribution.")
	wheelCmd.Flags().Bool("no-clean", false, "Don't clean up build directories.")
	wheelCmd.Flags().Bool("no-deps", false, "Don't install package dependencies.")
	wheelCmd.Flags().Bool("no-index", false, "Ignore package index.")
	wheelCmd.Flags().String("only-binary", "", "Do not use source packages.")
	wheelCmd.Flags().Bool("pre", false, "Include pre-release and development versions.")
	wheelCmd.Flags().Bool("prefer-binary", false, "Prefer older binary packages over newer source packages.")
	wheelCmd.Flags().String("progress-bar", "", "Specify type of progress to be displayed")
	wheelCmd.Flags().Bool("require-hashes", false, "Require a hash to check each requirement against")
	wheelCmd.Flags().StringP("requirement", "r", "", "Install from the given requirements file.")
	wheelCmd.Flags().String("src", "", "Directory to check out editable projects into.")
	wheelCmd.Flags().Bool("use-pep517", false, "Use PEP 517 for building source distributions.")
	wheelCmd.Flags().StringP("wheel-dir", "w", "", "Build wheels into <dir>")
	rootCmd.AddCommand(wheelCmd)

	carapace.Gen(wheelCmd).FlagCompletion(carapace.ActionMap{
		"constraint":   carapace.ActionFiles(),
		"editable":     carapace.ActionFiles(),
		"no-binary":    carapace.ActionValues(":all:", ":none:"),
		"only-binary":  carapace.ActionValues(":all:", ":none:"),
		"progress-bar": carapace.ActionValues("off", "on", "ascii", "pretty", "emoji"),
		"requirement":  carapace.ActionFiles(),
		"src":          carapace.ActionDirectories(),
		"wheel-dir":    carapace.ActionDirectories(),
	})

	carapace.Gen(wheelCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
