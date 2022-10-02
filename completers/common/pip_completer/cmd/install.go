package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pip"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().StringArray("abi", []string{}, "Only use wheels compatible with Python abi")
	installCmd.Flags().Bool("compile", false, "Compile Python source files to bytecode")
	installCmd.Flags().StringArrayP("constraint", "c", []string{}, "Constrain versions using the given constraints file")
	installCmd.Flags().StringP("editable", "e", "", "Install a project in editable mode")
	installCmd.Flags().String("extra-index-url", "", "Extra URLs of package indexes to use in addition to --index-url")
	installCmd.Flags().StringP("find-links", "f", "", "If a URL or path to an html file, then parse for links to archives such as sdist")
	installCmd.Flags().Bool("force-reinstall", false, "Reinstall all packages even if they are already up-to-date")
	installCmd.Flags().String("global-option", "", "Extra global options to be supplied to the setup.py call before the install command")
	installCmd.Flags().BoolP("ignore-installed", "I", false, "Ignore the installed packages, overwriting them")
	installCmd.Flags().Bool("ignore-requires-python", false, "Ignore the Requires-Python information")
	installCmd.Flags().String("implementation", "", "Only use wheels compatible with Python implementation <implementation>")
	installCmd.Flags().StringP("index-url", "i", "", "Base URL of the Python Package Index")
	installCmd.Flags().StringArray("install-option", []string{}, "Extra arguments to be supplied to the setup.py install command")
	installCmd.Flags().StringSlice("no-binary", []string{}, "Do not use binary packages")
	installCmd.Flags().Bool("no-build-isolation", false, "Disable isolation when building a modern source distribution")
	installCmd.Flags().Bool("no-clean", false, "Don't clean up build directories")
	installCmd.Flags().Bool("no-compile", false, "Do not compile Python source files to bytecode")
	installCmd.Flags().Bool("no-deps", false, "Don't install package dependencies")
	installCmd.Flags().Bool("no-index", false, "Ignore package index")
	installCmd.Flags().Bool("no-warn-conflicts", false, "Do not warn about broken dependencies")
	installCmd.Flags().Bool("no-warn-script-location", false, "Do not warn when installing scripts outside PATH")
	installCmd.Flags().StringArray("only-binary", []string{}, "Do not use source packages")
	installCmd.Flags().StringArray("platform", []string{}, "Only use wheels compatible with <platform>")
	installCmd.Flags().Bool("pre", false, "Include pre-release and development versions")
	installCmd.Flags().Bool("prefer-binary", false, "Prefer older binary packages over newer source packages")
	installCmd.Flags().String("prefix", "", "Installation prefix where lib, bin and other top-level folders are placed")
	installCmd.Flags().String("progress-bar", "", "Specify type of progress to be displayed")
	installCmd.Flags().String("python-version", "", "The Python interpreter version to use for wheel")
	installCmd.Flags().Bool("require-hashes", false, "Require a hash to check each requirement against")
	installCmd.Flags().StringArrayP("requirement", "r", []string{}, "Install from the given requirements file")
	installCmd.Flags().String("root", "", "Install everything relative to this alternate root directory")
	installCmd.Flags().String("src", "", "Directory to check out editable projects into")
	installCmd.Flags().StringP("target", "t", "", "Install packages into <dir>")
	installCmd.Flags().BoolP("upgrade", "U", false, "Upgrade all specified packages to the newest available version")
	installCmd.Flags().String("upgrade-strategy", "", "Determines how dependency upgrading should be handled")
	installCmd.Flags().Bool("use-pep517", false, "Use PEP 517 for building source distributions")
	installCmd.Flags().Bool("user", false, "Install to the Python user install directory for")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"constraint":       carapace.ActionFiles(),
		"editable":         carapace.ActionDirectories(),
		"implementation":   carapace.ActionValues("pp", "jy", "cp", "ip"),
		"progress-bar":     carapace.ActionValues("off", "on", "ascii", "pretty", "emoji"),
		"requirement":      carapace.ActionFiles(),
		"root":             carapace.ActionDirectories(),
		"src":              carapace.ActionDirectories(),
		"target":           carapace.ActionDirectories(),
		"upgrade-strategy": carapace.ActionValues("only-if-needed", "eager"),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			} else {
				// TODO support multiple index urls (--index-url) and update caching accordingly
				return pip.ActionPackageSearch()
			}
		}),
	)
}
