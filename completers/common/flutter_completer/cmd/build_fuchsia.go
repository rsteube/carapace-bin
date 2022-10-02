package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var build_fuchsiaCmd = &cobra.Command{
	Use:   "fuchsia",
	Short: "Build the Fuchsia target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_fuchsiaCmd).Standalone()

	build_fuchsiaCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	build_fuchsiaCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_fuchsiaCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_fuchsiaCmd.Flags().Bool("no-null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_fuchsiaCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_fuchsiaCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_fuchsiaCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_fuchsiaCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_fuchsiaCmd.Flags().String("runner-source", "", "The package source to use for the flutter_runner.")
	build_fuchsiaCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_fuchsiaCmd.Flags().String("target-platform", "", "The target platform for which the app is compiled.")
	build_fuchsiaCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_fuchsiaCmd)

	carapace.Gen(build_fuchsiaCmd).FlagCompletion(carapace.ActionMap{
		"runner-source": carapace.ActionValuesDescribed(
			"fuchsia.com", "runner already on the device",
			"flutter_tool", "runner distributed with Flutter",
		),
		"target":          carapace.ActionFiles(".dart"),
		"target-platform": carapace.ActionValues("fuchsia-arm64", "fuchsia-x64"),
	})
}
