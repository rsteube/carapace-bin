package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_iosFrameworkCmd = &cobra.Command{
	Use:   "ios-framework",
	Short: "Produces .xcframeworks for a Flutter project and its plugins for integration into existing, plain Xcode projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_iosFrameworkCmd).Standalone()

	build_iosFrameworkCmd.Flags().Bool("cocoapods", false, "Produce a Flutter.podspec instead of an engine Flutter.xcframework (recommended if host app uses CocoaPods).")
	build_iosFrameworkCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_iosFrameworkCmd.Flags().Bool("debug", false, "Produce a framework for the debug build configuration.")
	build_iosFrameworkCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_iosFrameworkCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_iosFrameworkCmd.Flags().Bool("no-cocoapods", false, "Do not produce a Flutter.podspec instead of an engine Flutter.xcframework (recommended if host app uses CocoaPods).")
	build_iosFrameworkCmd.Flags().Bool("no-debug", false, "Do not produce a framework for the debug build configuration.")
	build_iosFrameworkCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_iosFrameworkCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_iosFrameworkCmd.Flags().Bool("no-profile", false, "Do not produce a framework for the profile build configuration.")
	build_iosFrameworkCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_iosFrameworkCmd.Flags().Bool("no-release", false, "Do not produce a framework for the release build configuration.")
	build_iosFrameworkCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_iosFrameworkCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_iosFrameworkCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_iosFrameworkCmd.Flags().StringP("output", "o", "", "Location to write the frameworks.")
	build_iosFrameworkCmd.Flags().Bool("profile", false, "Produce a framework for the profile build configuration.")
	build_iosFrameworkCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_iosFrameworkCmd.Flags().Bool("release", false, "Produce a framework for the release build configuration.")
	build_iosFrameworkCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_iosFrameworkCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_iosFrameworkCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_iosFrameworkCmd)

	carapace.Gen(build_iosFrameworkCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionDirectories(),
		"target": carapace.ActionFiles(".dart"),
	})
}
