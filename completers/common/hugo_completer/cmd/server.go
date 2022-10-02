package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hugo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A high performance webserver",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	serverCmd.Flags().Bool("appendPort", true, "append port to baseURL")
	serverCmd.Flags().StringP("baseURL", "b", "", "hostname (and path) to the root, e.g. http://spf13.com/")
	serverCmd.Flags().String("bind", "127.0.0.1", "interface to which the server will bind")
	serverCmd.Flags().BoolP("buildDrafts", "D", false, "include content marked as draft")
	serverCmd.Flags().BoolP("buildExpired", "E", false, "include expired content")
	serverCmd.Flags().BoolP("buildFuture", "F", false, "include content with publishdate in the future")
	serverCmd.Flags().String("cacheDir", "", "filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/")
	serverCmd.Flags().Bool("cleanDestinationDir", false, "remove files from destination not found in static directories")
	serverCmd.Flags().StringP("contentDir", "c", "", "filesystem path to content directory")
	serverCmd.Flags().StringP("destination", "d", "", "filesystem path to write files to")
	serverCmd.Flags().Bool("disableBrowserError", false, "do not show build errors in the browser")
	serverCmd.Flags().Bool("disableFastRender", false, "enables full re-renders on changes")
	serverCmd.Flags().StringSlice("disableKinds", []string{}, "disable different kind of pages (home, RSS etc.)")
	serverCmd.Flags().Bool("disableLiveReload", false, "watch without enabling live browser reload on rebuild")
	serverCmd.Flags().Bool("enableGitInfo", false, "add Git revision, date and author info to the pages")
	serverCmd.Flags().Bool("forceSyncStatic", false, "copy all files when static is changed.")
	serverCmd.Flags().Bool("gc", false, "enable to run some cleanup tasks (remove unused cache files) after the build")
	serverCmd.Flags().Bool("i18n-warnings", false, "print missing translations")
	serverCmd.Flags().Bool("ignoreCache", false, "ignores the cache directory")
	serverCmd.Flags().StringP("layoutDir", "l", "", "filesystem path to layout directory")
	serverCmd.Flags().Int("liveReloadPort", -1, "port for live reloading (i.e. 443 in HTTPS proxy situations)")
	serverCmd.Flags().String("meminterval", "100ms", "interval to poll memory usage (requires --memstats), valid time units are \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\".")
	serverCmd.Flags().String("memstats", "", "log memory usage to this file")
	serverCmd.Flags().Bool("minify", false, "minify any supported output format (HTML, XML etc.)")
	serverCmd.Flags().Bool("navigateToChanged", false, "navigate to changed content file on live browser reload")
	serverCmd.Flags().Bool("noChmod", false, "don't sync permission mode of files")
	serverCmd.Flags().Bool("noHTTPCache", false, "prevent HTTP caching")
	serverCmd.Flags().Bool("noTimes", false, "don't sync modification time of files")
	serverCmd.Flags().Bool("path-warnings", false, "print warnings on duplicate target paths etc.")
	serverCmd.Flags().String("poll", "", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes")
	serverCmd.Flags().IntP("port", "p", 1313, "port on which the server will listen")
	serverCmd.Flags().Bool("print-mem", false, "print memory usage to screen at intervals")
	serverCmd.Flags().String("profile-cpu", "", "write cpu profile to `file`")
	serverCmd.Flags().String("profile-mem", "", "write memory profile to `file`")
	serverCmd.Flags().String("profile-mutex", "", "write Mutex profile to `file`")
	serverCmd.Flags().Bool("renderToDisk", false, "render to Destination path (default is render to memory & serve from there)")
	serverCmd.Flags().Bool("templateMetrics", false, "display metrics about template executions")
	serverCmd.Flags().Bool("templateMetricsHints", false, "calculate some improvement hints when combined with --templateMetrics")
	serverCmd.Flags().StringSliceP("theme", "t", []string{}, "themes to use (located in /themes/THEMENAME/)")
	serverCmd.Flags().String("trace", "", "write trace to `file` (not useful in general)")
	serverCmd.Flags().BoolP("watch", "w", true, "watch filesystem for changes and recreate as needed")
	rootCmd.AddCommand(serverCmd)

	carapace.Gen(serverCmd).FlagCompletion(carapace.ActionMap{
		"cacheDir":      carapace.ActionDirectories(),
		"contentDir":    carapace.ActionDirectories(),
		"destination":   carapace.ActionDirectories(),
		"disableKinds":  action.ActionKinds(),
		"layoutDir":     carapace.ActionDirectories(),
		"profile-cpu":   carapace.ActionFiles(),
		"profile-mem":   carapace.ActionFiles(),
		"profile-mutex": carapace.ActionFiles(),
		"theme":         action.ActionThemes(serverCmd),
		"trace":         carapace.ActionFiles(),
	})
}
