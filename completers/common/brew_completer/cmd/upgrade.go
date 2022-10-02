package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade outdated casks and outdated, unpinned formulae",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("appdir", "", "Target location for Applications")
	upgradeCmd.Flags().String("audio-unit-plugindir", "", "Target location for Audio Unit Plugins")
	upgradeCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables")
	upgradeCmd.Flags().BoolP("build-from-source", "s", false, "Compile formula from source even if a bottle is available")
	upgradeCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	upgradeCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	upgradeCmd.Flags().String("colorpickerdir", "", "Target location for Color Pickers")
	upgradeCmd.Flags().BoolP("debug", "d", false, "If brewing fails, open an interactive debugging session")
	upgradeCmd.Flags().String("dictionarydir", "", "Target location for Dictionaries")
	upgradeCmd.Flags().Bool("display-times", false, "Print install times for each package at the end of the run")
	upgradeCmd.Flags().BoolP("dry-run", "n", false, "Show what would be upgraded")
	upgradeCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation is outdated")
	upgradeCmd.Flags().String("fontdir", "", "Target location for Fonts")
	upgradeCmd.Flags().BoolP("force", "f", false, "Install formulae without checking for previously installed versions")
	upgradeCmd.Flags().Bool("force-bottle", false, "Install from a bottle even if it would not normally be used")
	upgradeCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	upgradeCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	upgradeCmd.Flags().Bool("greedy", false, "Also include casks with auto_updates true")
	upgradeCmd.Flags().Bool("greedy-auto-updates", false, "Also include casks with auto_updates true.")
	upgradeCmd.Flags().Bool("greedy-latest", false, "Also include casks with version :latest.")
	upgradeCmd.Flags().BoolP("help", "h", false, "Show this message.")
	upgradeCmd.Flags().Bool("ignore-pinned", false, "Set a successful exit status even if pinned formulae are not upgraded")
	upgradeCmd.Flags().String("input-methoddir", "", "Target location for Input Methods")
	upgradeCmd.Flags().BoolP("interactive", "i", false, "Download and patch formula")
	upgradeCmd.Flags().String("internet-plugindir", "", "Target location for Internet Plugins")
	upgradeCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation")
	upgradeCmd.Flags().Bool("language", false, "Comma-separated list of language codes to prefer")
	upgradeCmd.Flags().String("mdimporterdir", "", "Target location for Spotlight Plugins")
	upgradeCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables")
	upgradeCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads")
	upgradeCmd.Flags().String("prefpanedir", "", "Target location for Preference Panes")
	upgradeCmd.Flags().String("qlplugindir", "", "Target location for QuickLook Plugins")
	upgradeCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads")
	upgradeCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	upgradeCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	upgradeCmd.Flags().String("screen-saverdir", "", "Target location for Screen Savers")
	upgradeCmd.Flags().String("servicedir", "", "Target location for Services")
	upgradeCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	upgradeCmd.Flags().BoolP("verbose", "v", false, "Print the verification and postinstall steps")
	upgradeCmd.Flags().String("vst-plugindir", "", "Target location for VST Plugins")
	upgradeCmd.Flags().String("vst3-plugindir", "", "Target location for VST3 Plugins")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionList(upgradeCmd),
	)
}
