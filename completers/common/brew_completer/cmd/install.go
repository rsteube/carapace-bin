package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a formula or cask",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("HEAD", false, "If formula defines it, install the HEAD version")
	installCmd.Flags().String("appdir", "", "Target location for Applications")
	installCmd.Flags().String("audio-unit-plugindir", "", "Target location for Audio Unit Plugins")
	installCmd.Flags().Bool("binaries", false, "Enable linking of helper executables")
	installCmd.Flags().String("bottle-arch", "", "Optimise bottles for the specified architecture")
	installCmd.Flags().Bool("build-bottle", false, "Prepare the formula for eventual bottling during installation")
	installCmd.Flags().BoolP("build-from-source", "s", false, "Compile formula from source even if a bottle is provided")
	installCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	installCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	installCmd.Flags().Bool("cc", false, "Attempt to compile using the specified compiler")
	installCmd.Flags().String("colorpickerdir", "", "Target location for Color Pickers")
	installCmd.Flags().BoolP("debug", "d", false, "If brewing fails, open an interactive debugging session")
	installCmd.Flags().String("dictionarydir", "", "Target location for Dictionaries (default:")
	installCmd.Flags().Bool("display-times", false, "Print install times for each package at the end of the run")
	installCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation of the formula is outdated")
	installCmd.Flags().String("fontdir", "", "Target location for Fonts (default:")
	installCmd.Flags().BoolP("force", "f", false, "Install formulae without checking for previously installed")
	installCmd.Flags().Bool("force-bottle", false, "Install from a bottle even if it would not normally be used for installation")
	installCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	installCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	installCmd.Flags().BoolP("git", "g", false, "Create a Git repository")
	installCmd.Flags().BoolP("help", "h", false, "Show this message.")
	installCmd.Flags().Bool("ignore-dependencies", false, "An unsupported Homebrew development flag to skip installing any dependencies")
	installCmd.Flags().Bool("include-test", false, "Install testing dependencies required to run brew test formula")
	installCmd.Flags().String("input-methoddir", "", "Target location for Input Methods (default:")
	installCmd.Flags().BoolP("interactive", "i", false, "Download and patch formula, then open a shell")
	installCmd.Flags().String("internet-plugindir", "", "Target location for Internet Plugins")
	installCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation")
	installCmd.Flags().String("language", "", "Comma-separated list of language codes to prefer for cask installation")
	installCmd.Flags().String("mdimporterdir", "", "Target location for Spotlight Plugins")
	installCmd.Flags().Bool("no-binaries", false, "Disable linking of helper executables")
	installCmd.Flags().Bool("no-quarantine", false, "Disable quarantining of downloads")
	installCmd.Flags().Bool("only-dependencies", false, "Install the dependencies with specified options")
	installCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking")
	installCmd.Flags().String("prefpanedir", "", "Target location for Preference Panes")
	installCmd.Flags().String("qlplugindir", "", "Target location for QuickLook Plugins")
	installCmd.Flags().Bool("quarantine", false, "Enable quarantining of downloads")
	installCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	installCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	installCmd.Flags().String("screen-saverdir", "", "Target location for Screen Savers (default:")
	installCmd.Flags().String("servicedir", "", "Target location for Services (default:")
	installCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies")
	installCmd.Flags().BoolP("verbose", "v", false, "Print the verification and postinstall steps")
	installCmd.Flags().String("vst-plugindir", "", "Target location for VST Plugins (default:")
	installCmd.Flags().String("vst3-plugindir", "", "Target location for VST3 Plugins (default:")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"appdir":               carapace.ActionDirectories(),
		"audio-unit-plugindir": carapace.ActionDirectories(),
		"colorpickerdir":       carapace.ActionDirectories(),
		"dictionarydir":        carapace.ActionDirectories(),
		"fontdir":              carapace.ActionDirectories(),
		"input-methoddir":      carapace.ActionDirectories(),
		"internet-plugindir":   carapace.ActionDirectories(),
		"mdimporterdir":        carapace.ActionDirectories(),
		"prefpanedir":          carapace.ActionDirectories(),
		"qlplugindir":          carapace.ActionDirectories(),
		"screen-saverdir":      carapace.ActionDirectories(),
		"servicedir":           carapace.ActionDirectories(),
		"vst-plugindir":        carapace.ActionDirectories(),
		"vst3-plugindir":       carapace.ActionDirectories(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionSearch(installCmd),
	)
}
