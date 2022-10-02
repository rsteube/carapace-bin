package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_webCmd = &cobra.Command{
	Use:   "web",
	Short: "Build a web application bundle",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_webCmd).Standalone()

	build_webCmd.Flags().Bool("csp", false, "Disable dynamic generation of code in the generated output.")
	build_webCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_webCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_webCmd.Flags().Bool("native-null-assertions", false, "Enables additional runtime null checks in web applications.")
	build_webCmd.Flags().Bool("no-native-null-assertions", false, "Disables additional runtime null checks in web applications.")
	build_webCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_webCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_webCmd.Flags().Bool("no-source-maps", false, "Do not generate a sourcemap file.")
	build_webCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_webCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_webCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_webCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_webCmd.Flags().Bool("pwa-strategy", false, "val                    The caching strategy to be used by the PWA service worker.")
	build_webCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_webCmd.Flags().Bool("source-maps", false, "Generate a sourcemap file.")
	build_webCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_webCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	build_webCmd.Flags().String("web-renderer", "", "The renderer implementation to use when building for the web.")
	buildCmd.AddCommand(build_webCmd)

	carapace.Gen(build_webCmd).FlagCompletion(carapace.ActionMap{
		"pwa-strategy": carapace.ActionValuesDescribed(
			"none", "Generate a service worker with no body.",
			"offline-first", "Attempt to cache the application shell eagerly.",
		),
		"target": carapace.ActionFiles(".dart"),
		"web-renderer": carapace.ActionValuesDescribed(
			"auto", "Use the HTML renderer on mobile devices, and CanvasKit on desktop devices.",
			"canvaskit", "Always use the CanvasKit renderer.",
			"html", "Always use the HTML renderer.",
		),
	})

}
