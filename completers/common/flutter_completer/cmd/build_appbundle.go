package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_appbundleCmd = &cobra.Command{
	Use:   "appbundle",
	Short: "Build an Android App Bundle file from your app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_appbundleCmd).Standalone()

	build_appbundleCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_appbundleCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users.")
	build_appbundleCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_appbundleCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_appbundleCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_appbundleCmd.Flags().Bool("deferred-components", false, "Setting to false disables building with deferred components.")
	build_appbundleCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_appbundleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_appbundleCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_appbundleCmd.Flags().Bool("no-deferred-components", false, "Setting to false disables building with deferred components.")
	build_appbundleCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_appbundleCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_appbundleCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_appbundleCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_appbundleCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_appbundleCmd.Flags().Bool("no-validate-deferred-components", false, "When enabled, deferred component apps will fail to build if setup problems are detected that would prevent deferred components from functioning properly.")
	build_appbundleCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_appbundleCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_appbundleCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_appbundleCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_appbundleCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_appbundleCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_appbundleCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_appbundleCmd.Flags().String("target-platform", "", "The target platform for which the app is compiled.")
	build_appbundleCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_appbundleCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	build_appbundleCmd.Flags().Bool("validate-deferred-components", false, "When enabled, deferred component apps will fail to build if setup problems are detected that would prevent deferred components from functioning properly.")
	buildCmd.AddCommand(build_appbundleCmd)

	carapace.Gen(build_appbundleCmd).FlagCompletion(carapace.ActionMap{
		"target":          carapace.ActionFiles(".dart"),
		"target-platform": carapace.ActionValues("android-arm", "android-arm64", "android-x64"),
	})
}
