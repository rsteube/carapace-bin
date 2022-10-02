package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var driveCmd = &cobra.Command{
	Use:   "drive",
	Short: "Run integration tests for the project on an attached device or emulator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(driveCmd).Standalone()

	driveCmd.Flags().Bool("android-emulator", false, "Perform Flutter Driver testing using an Android Emulator.")
	driveCmd.Flags().StringP("android-project-arg", "P", "", "Additional arguments specified as key=value that are passed directly to the gradle project via the -P flag.")
	driveCmd.Flags().String("browser-dimension", "", "The dimension of the browser when running a Flutter Web test.")
	driveCmd.Flags().String("browser-name", "", "Name of the browser where tests will be executed.")
	driveCmd.Flags().Bool("build", false, "(deprecated) Build the app before running.")
	driveCmd.Flags().Bool("cache-sksl", false, "Cache the shader in the SkSL format instead of in binary or GLSL formats.")
	driveCmd.Flags().String("chrome-binary", "", "Location of the Chrome binary.")
	driveCmd.Flags().StringArray("dart-define", []string{}, "Additional key-value pairs that will be available as constants.")
	driveCmd.Flags().StringP("dart-entrypoint-args", "a", "", "Pass a list of arguments to the Dart entrypoint at application startup.")
	driveCmd.Flags().String("dds-port", "", "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port.")
	driveCmd.Flags().Bool("debug", false, "Build a debug version of your app (default mode).")
	driveCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach.")
	driveCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only.")
	driveCmd.Flags().String("device-vmservice-port", "", "Look for vmservice connections only from the specified port.")
	driveCmd.Flags().String("driver", "", "The test file to run on the host (as opposed to the target file to run on the device).")
	driveCmd.Flags().String("driver-port", "", "The port where Webdriver server is launched at.")
	driveCmd.Flags().Bool("dump-skp-on-shader-compilation", false, "Automatically dump the skp that triggers new shader compilations.")
	driveCmd.Flags().Bool("endless-trace-buffer", false, "Enable tracing to an infinite buffer, instead of a ring buffer.")
	driveCmd.Flags().Bool("flavor", false, "Build a custom app flavor as defined by platform-specific build setup.")
	driveCmd.Flags().Bool("headless", false, "Launch driver browser in headless mode.")
	driveCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	driveCmd.Flags().String("host-vmservice-port", "", "When a device-side vmservice port is forwarded to a host-side port, use this value as the host port.")
	driveCmd.Flags().Bool("keep-app-running", false, "Will keep the Flutter application running when done testing.")
	driveCmd.Flags().Bool("no-android-emulator", false, "Do not perform Flutter Driver testing using an Android Emulator.")
	driveCmd.Flags().Bool("no-build", false, "(deprecated) Do not build the app before running.")
	driveCmd.Flags().Bool("no-headless", false, "Do not launch driver browser in headless mode.")
	driveCmd.Flags().Bool("no-keep-app-running", false, "Will not keep the Flutter application running when done testing.")
	driveCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	driveCmd.Flags().Bool("no-start-paused", false, "Do not start in a paused mode and wait for a debugger to connect.")
	driveCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	driveCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	driveCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	driveCmd.Flags().String("profile-memory", "", "Launch devtools and profile application memory, writing The output data to the file path provided to this argument as JSON.")
	driveCmd.Flags().Bool("purge-persistent-cache", false, "Removes all existing persistent caches. This allows reproducing shader compilation jank that normally only happens the first time an app is run, or for reliable testing of compilation jank fixes (e.g. shader warm-up).")
	driveCmd.Flags().Bool("release", false, "Build a release version of your app.")
	driveCmd.Flags().String("route", "", "Which route to load when running the app.")
	driveCmd.Flags().String("screenshot", "", "Directory location to write screenshots on test failure.")
	driveCmd.Flags().Bool("start-paused", false, "Start in a paused mode and wait for a debugger to connect.")
	driveCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	driveCmd.Flags().String("test-arguments", "", "Additional arguments to pass to the Dart VM running The test script.")
	driveCmd.Flags().Bool("trace-skia", false, "Enable tracing of Skia code.")
	driveCmd.Flags().Bool("trace-startup", false, "Trace application startup, then exit, saving the trace to a file.")
	driveCmd.Flags().Bool("trace-systrace", false, "Enable tracing to the system tracer.")
	driveCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	driveCmd.Flags().String("use-application-binary", "", "Specify a pre-built application binary to use when running.")
	driveCmd.Flags().String("use-existing-app", "", "Connect to an already running instance via the given observatory URL.")
	driveCmd.Flags().Bool("verbose-system-logs", false, "Include verbose logging from the Flutter engine.")
	driveCmd.Flags().String("web-renderer", "", "The renderer implementation to use when building for the web.")
	driveCmd.Flags().Bool("write-sksl-on-exit", false, "Attempts to write an SkSL file when the drive process is finished to the provided file, overwriting it if necessary.")
	rootCmd.AddCommand(driveCmd)

	carapace.Gen(driveCmd).FlagCompletion(carapace.ActionMap{
		"browser-name": carapace.ActionValuesDescribed(
			"android-chrome", "Chrome on Android.",
			"chrome", "Google Chrome on this computer.",
			"edge", "Microsoft Edge on this computer (Windows only).",
			"firefox", "Mozilla Firefox on this computer.",
			"ios-safari", "Apple Safari on an iOS device.",
			"safari", "Apple Safari on this computer (macOS only).",
		),
		"chrome-binary":  carapace.ActionFiles(),
		"device-user":    adb.ActionUsers(),
		"driver":         carapace.ActionFiles(),
		"profile-memory": carapace.ActionFiles(),
		"screenshot":     carapace.ActionDirectories(),
		"target":         carapace.ActionFiles(".dart"),
		"web-renderer": carapace.ActionValuesDescribed(
			"auto", "Use the HTML renderer on mobile devices, and CanvasKit on desktop devices.",
			"canvaskit", "Always use the CanvasKit renderer.",
			"html", "Always use the HTML renderer.",
		),
	})
}
