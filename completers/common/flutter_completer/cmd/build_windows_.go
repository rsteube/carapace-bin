package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "Build a Windows desktop application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_windowsCmd).Standalone()

	build_windowsCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_windowsCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_windowsCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_windowsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_windowsCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_windowsCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_windowsCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_windowsCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_windowsCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_windowsCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_windowsCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_windowsCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_windowsCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_windowsCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_windowsCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_windowsCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_windowsCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_windowsCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_windowsCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_windowsCmd)

	carapace.Gen(build_windowsCmd).FlagCompletion(carapace.ActionMap{
		"target": carapace.ActionFiles(".dart"),
	})
}
