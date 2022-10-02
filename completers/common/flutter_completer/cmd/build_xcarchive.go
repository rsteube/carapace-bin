package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_xcarchiveCmd = &cobra.Command{
	Use:   "xcarchive",
	Short: "Build an iOS archive bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_xcarchiveCmd).Standalone()

	build_xcarchiveCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_xcarchiveCmd.Flags().String("build-name", "", "A \"x.y.z\" string used as the version number shown to users.")
	build_xcarchiveCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_xcarchiveCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_xcarchiveCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_xcarchiveCmd.Flags().String("export-options-plist", "", "Optionally export an IPA with these options.")
	build_xcarchiveCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	build_xcarchiveCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_xcarchiveCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_xcarchiveCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_xcarchiveCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_xcarchiveCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_xcarchiveCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_xcarchiveCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_xcarchiveCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_xcarchiveCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_xcarchiveCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_xcarchiveCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_xcarchiveCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_xcarchiveCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_xcarchiveCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_xcarchiveCmd)

	carapace.Gen(build_xcarchiveCmd).FlagCompletion(carapace.ActionMap{
		"export-options-plist": carapace.ActionFiles(),
		"target":               carapace.ActionFiles(".dart"),
	})
}
