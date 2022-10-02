package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_iosCmd = &cobra.Command{
	Use:   "ios",
	Short: "Build an iOS application bundle (Mac OS X host only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_iosCmd).Standalone()

	build_iosCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_iosCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users.")
	build_iosCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_iosCmd.Flags().Bool("codesign", false, "Codesign the application bundle (only available on device builds).")
	build_iosCmd.Flags().Bool("config-only", false, "Update the project configuration without performing a build.")
	build_iosCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_iosCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_iosCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_iosCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_iosCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_iosCmd.Flags().Bool("no-codesign", false, "Codesign the application bundle (only available on device builds).")
	build_iosCmd.Flags().Bool("no-config-only", false, "Update the project configuration without performing a build.")
	build_iosCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_iosCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_iosCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_iosCmd.Flags().Bool("no-simulator", false, "Build for the iOS simulator instead of the device.")
	build_iosCmd.Flags().Bool("no-tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	build_iosCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_iosCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_iosCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_iosCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_iosCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_iosCmd.Flags().Bool("simulator", false, "Build for the iOS simulator instead of the device.")
	build_iosCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_iosCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_iosCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_iosCmd)

	carapace.Gen(build_iosCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionFiles(".dart"),
	})
}
