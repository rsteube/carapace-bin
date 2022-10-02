package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_aarCmd = &cobra.Command{
	Use:   "aar",
	Short: "Build a repository containing an AAR and a POM file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_aarCmd).Standalone()

	build_aarCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_aarCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_aarCmd.Flags().Bool("debug", false, "Build a debug version of the current project.")
	build_aarCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_aarCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_aarCmd.Flags().Bool("no-debug", false, "Do not build a debug version of the current project.")
	build_aarCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_aarCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_aarCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_aarCmd.Flags().Bool("no-release", false, "Do not build a release version of the current project.")
	build_aarCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_aarCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_aarCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_aarCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_aarCmd.Flags().String("output-dir", "", "The absolute path to the directory where the repository is generated.")
	build_aarCmd.Flags().Bool("profile", false, "Build a version of the current project specialized for performance profiling.")
	build_aarCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_aarCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_aarCmd.Flags().String("target-platform", "", "The target platform for which the project is compiled.")
	build_aarCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_aarCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_aarCmd)

	carapace.Gen(build_aarCmd).FlagCompletion(carapace.ActionMap{
		"output-dir":      carapace.ActionDirectories(),
		"target-platform": carapace.ActionValues("android-arm", "android-arm64", "android-x86", "android-x64"),
	})
}
