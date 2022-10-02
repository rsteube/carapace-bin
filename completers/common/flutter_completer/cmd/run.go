package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run your Flutter app on an attached device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("android-project-arg", "P", "", "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag.")
	runCmd.Flags().Bool("await-first-frame-when-tracing", false, "Wait for the first frame when tracing startup.")
	runCmd.Flags().Bool("build", false, "If necessary, build the app before running.")
	runCmd.Flags().Bool("cache-sksl", false, "Cache the shader in the SkSL format instead of in binary or GLSL formats.")
	runCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	runCmd.Flags().StringP("dart-entrypoint-args", "a", "", "Pass a list of arguments to the Dart entrypoint at application startup.")
	runCmd.Flags().String("dds-port", "", "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port.")
	runCmd.Flags().Bool("debug", false, "Build a debug version of your app (default mode).")
	runCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	runCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only.")
	runCmd.Flags().String("device-vmservice-port", "", "Look for vmservice connections only from the specified port.")
	runCmd.Flags().Bool("dump-skp-on-shader-compilation", false, "Automatically dump the skp that triggers new shader compilations.")
	runCmd.Flags().Bool("enable-software-rendering", false, "Enable rendering using the Skia software backend.")
	runCmd.Flags().Bool("endless-trace-buffer", false, "Enable tracing to an infinite buffer, instead of a ring buffer.")
	runCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	runCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	runCmd.Flags().String("host-vmservice-port", "", "When a device-side vmservice port is forwarded to a host-side port, use this value as the host port.")
	runCmd.Flags().Bool("hot", false, "Run with support for hot reloading.")
	runCmd.Flags().Bool("no-await-first-frame-when-tracing", false, "Do not wait for the first frame when tracing startup.")
	runCmd.Flags().Bool("no-build", false, "If necessary, do not build the app before running.")
	runCmd.Flags().Bool("no-hot", false, "Run without support for hot reloading.")
	runCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	runCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	runCmd.Flags().Bool("no-start-paused", false, "Do not start in a paused mode and wait for a debugger to connect.")
	runCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	runCmd.Flags().Bool("no-use-test-fonts", false, "Disable the \"Ahem\" font.")
	runCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	runCmd.Flags().String("pid-file", "", "Specify a file to write the process ID to.")
	runCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	runCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	runCmd.Flags().Bool("purge-persistent-cache", false, "Removes all existing persistent caches.")
	runCmd.Flags().Bool("release", false, "Build a release version of your app.")
	runCmd.Flags().Bool("route", false, "Which route to load when running the app.")
	runCmd.Flags().Bool("skia-deterministic-rendering", false, "Provide completely deterministic (i.e. reproducible) Skia rendering.")
	runCmd.Flags().Bool("start-paused", false, "Start in a paused mode and wait for a debugger to connect.")
	runCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	runCmd.Flags().Bool("trace-skia", false, "Enable tracing of Skia code.")
	runCmd.Flags().Bool("trace-startup", false, "Trace application startup, then exit, saving the trace to a file.")
	runCmd.Flags().Bool("trace-systrace", false, "Enable tracing to the system tracer.")
	runCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	runCmd.Flags().String("use-application-binary", "", "Specify a pre-built application binary to use when running.")
	runCmd.Flags().Bool("use-test-fonts", false, "Enable the \"Ahem\" font.")
	runCmd.Flags().Bool("verbose-system-logs", false, "Include verbose logging from the Flutter engine.")
	runCmd.Flags().Bool("web-renderer", false, "The renderer implementation to use when building for the web.")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"device-user":            adb.ActionUsers(),
		"pid-file":               carapace.ActionFiles(),
		"target":                 carapace.ActionFiles(".dart"),
		"use-application-binary": carapace.ActionFiles(".apk", ".ipa"),
	})
}
