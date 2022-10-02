package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hugo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "hugo builds your site",
	Long:  "https://gohugo.io/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringP("baseURL", "b", "", "hostname (and path) to the root, e.g. http://spf13.com/")
	rootCmd.Flags().BoolP("buildDrafts", "D", false, "include content marked as draft")
	rootCmd.Flags().BoolP("buildExpired", "E", false, "include expired content")
	rootCmd.Flags().BoolP("buildFuture", "F", false, "include content with publishdate in the future")
	rootCmd.Flags().String("cacheDir", "", "filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/")
	rootCmd.Flags().Bool("cleanDestinationDir", false, "remove files from destination not found in static directories")
	rootCmd.PersistentFlags().String("config", "", "config file (default is path/config.yaml|json|toml)")
	rootCmd.PersistentFlags().String("configDir", "config", "config dir")
	rootCmd.Flags().StringP("contentDir", "c", "", "filesystem path to content directory")
	rootCmd.PersistentFlags().Bool("debug", false, "debug output")
	rootCmd.Flags().StringP("destination", "d", "", "filesystem path to write files to")
	rootCmd.Flags().StringSlice("disableKinds", []string{}, "disable different kind of pages (home, RSS etc.)")
	rootCmd.Flags().Bool("enableGitInfo", false, "add Git revision, date and author info to the pages")
	rootCmd.PersistentFlags().StringP("environment", "e", "", "build environment")
	rootCmd.Flags().Bool("forceSyncStatic", false, "copy all files when static is changed.")
	rootCmd.Flags().Bool("gc", false, "enable to run some cleanup tasks (remove unused cache files) after the build")
	rootCmd.Flags().Bool("i18n-warnings", false, "print missing translations")
	rootCmd.Flags().Bool("ignoreCache", false, "ignores the cache directory")
	rootCmd.PersistentFlags().Bool("ignoreVendor", false, "ignores any _vendor directory")
	rootCmd.PersistentFlags().String("ignoreVendorPaths", "", "ignores any _vendor for module paths matching the given Glob pattern")
	rootCmd.Flags().StringP("layoutDir", "l", "", "filesystem path to layout directory")
	rootCmd.PersistentFlags().Bool("log", false, "enable Logging")
	rootCmd.PersistentFlags().String("logFile", "", "log File path (if set, logging enabled automatically)")
	rootCmd.Flags().Bool("minify", false, "minify any supported output format (HTML, XML etc.)")
	rootCmd.Flags().Bool("noChmod", false, "don't sync permission mode of files")
	rootCmd.Flags().Bool("noTimes", false, "don't sync modification time of files")
	rootCmd.Flags().Bool("path-warnings", false, "print warnings on duplicate target paths etc.")
	rootCmd.Flags().String("poll", "", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes")
	rootCmd.Flags().Bool("print-mem", false, "print memory usage to screen at intervals")
	rootCmd.Flags().String("profile-cpu", "", "write cpu profile to `file`")
	rootCmd.Flags().String("profile-mem", "", "write memory profile to `file`")
	rootCmd.Flags().String("profile-mutex", "", "write Mutex profile to `file`")
	rootCmd.PersistentFlags().Bool("quiet", false, "build in quiet mode")
	rootCmd.Flags().Bool("renderToMemory", false, "render to memory (only useful for benchmark testing)")
	rootCmd.PersistentFlags().StringP("source", "s", "", "filesystem path to read files relative from")
	rootCmd.Flags().Bool("templateMetrics", false, "display metrics about template executions")
	rootCmd.Flags().Bool("templateMetricsHints", false, "calculate some improvement hints when combined with --templateMetrics")
	rootCmd.Flags().StringSliceP("theme", "t", []string{}, "themes to use (located in /themes/THEMENAME/)")
	rootCmd.PersistentFlags().String("themesDir", "", "filesystem path to themes directory")
	rootCmd.Flags().String("trace", "", "write trace to `file` (not useful in general)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().Bool("verboseLog", false, "verbose logging")
	rootCmd.Flags().BoolP("watch", "w", false, "watch filesystem for changes and recreate as needed")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cacheDir":      carapace.ActionDirectories(),
		"config":        carapace.ActionFiles(".yaml", ".yml", ".json", ".toml"),
		"configDir":     carapace.ActionDirectories(),
		"contentDir":    carapace.ActionDirectories(),
		"destination":   carapace.ActionDirectories(),
		"disableKinds":  action.ActionKinds(),
		"environment":   action.ActionEnvironments(rootCmd),
		"layoutDir":     carapace.ActionDirectories(),
		"logFile":       carapace.ActionFiles(),
		"profile-cpu":   carapace.ActionFiles(),
		"profile-mem":   carapace.ActionFiles(),
		"profile-mutex": carapace.ActionFiles(),
		"source":        carapace.ActionDirectories(),
		"theme":         action.ActionThemes(rootCmd),
		"themesDir":     carapace.ActionDirectories(),
		"trace":         carapace.ActionFiles(),
	})
}
