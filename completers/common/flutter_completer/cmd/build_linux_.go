package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_linuxCmd = &cobra.Command{
	Use:   "linux",
	Short: "Build a Linux desktop application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_linuxCmd).Standalone()

	build_linuxCmd.Flags().Bool("analyze-size", false, "Whether to produce additional profile information for artifact output size.")
	build_linuxCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_linuxCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_linuxCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_linuxCmd.Flags().Bool("no-analyze-size", false, "Whether to produce additional profile information for artifact output size.")
	build_linuxCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_linuxCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values.")
	build_linuxCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_linuxCmd.Flags().Bool("no-track-widget-creation", false, "Do no track widget creation locations.")
	build_linuxCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_linuxCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_linuxCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values.")
	build_linuxCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_linuxCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_linuxCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_linuxCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_linuxCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_linuxCmd.Flags().String("target-platform", "", "The target platform for which the app is compiled.")
	build_linuxCmd.Flags().String("target-sysroot", "", "The root filesystem path of target platform for which the app is compiled.")
	build_linuxCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_linuxCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_linuxCmd)

	carapace.Gen(build_linuxCmd).FlagCompletion(carapace.ActionMap{
		"target":          carapace.ActionFiles(".dart"),
		"target-platform": carapace.ActionValues("linux-arm64", "linux-x64"),
	})
}
