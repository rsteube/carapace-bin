package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_apkCmd = &cobra.Command{
	Use:   "apk",
	Short: "Build an Android APK file from your app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_apkCmd).Standalone()

	build_apkCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_apkCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users.")
	build_apkCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_apkCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_apkCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_apkCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_apkCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_apkCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_apkCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_apkCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_apkCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_apkCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_apkCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_apkCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_apkCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_apkCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_apkCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_apkCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_apkCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_apkCmd.Flags().Bool("split-per-abi", false, "Whether to split the APKs per ABIs.")
	build_apkCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_apkCmd.Flags().String("target-platform", "", "The target platform for which the app is compiled.")
	build_apkCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_apkCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_apkCmd)

	carapace.Gen(build_apkCmd).FlagCompletion(carapace.ActionMap{
		"target":          carapace.ActionFiles(".dart"),
		"target-platform": carapace.ActionValues("android-arm", "android-arm64", "android-x86", "android-x64"),
	})
}
