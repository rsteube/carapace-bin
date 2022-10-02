package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Build the Flutter assets directory from your app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_bundleCmd).Standalone()

	build_bundleCmd.Flags().String("asset-dir", "", "The output directory for the kernel_blob.bin file, the native snapshet, the assets, etc.")
	build_bundleCmd.Flags().String("build-number", "", "An identifier used as an internal version number.")
	build_bundleCmd.Flags().Bool("debug", false, "Build a debug version of your app (default mode).")
	build_bundleCmd.Flags().String("depfile", "", "A file path where a depfile will be written.")
	build_bundleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_bundleCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_bundleCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_bundleCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_bundleCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_bundleCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_bundleCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_bundleCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_bundleCmd.Flags().String("target-platform", "", "The architecture for which to build the application.")
	build_bundleCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_bundleCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_bundleCmd)

	carapace.Gen(build_bundleCmd).FlagCompletion(carapace.ActionMap{
		"asset-dir":       carapace.ActionDirectories(),
		"depfile":         carapace.ActionFiles(),
		"target":          carapace.ActionFiles(".dart"),
		"target-platform": carapace.ActionValues("android-arm", "android-arm64", "android-x86", "android-x64"),
	})
}
