package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hugo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var modCmd = &cobra.Command{
	Use:   "mod",
	Short: "Various Hugo Modules helpers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	modCmd.Flags().StringP("baseURL", "b", "", "hostname (and path) to the root, e.g. http://spf13.com/")
	modCmd.Flags().BoolP("buildDrafts", "D", false, "include content marked as draft")
	modCmd.Flags().BoolP("buildExpired", "E", false, "include expired content")
	modCmd.Flags().BoolP("buildFuture", "F", false, "include content with publishdate in the future")
	modCmd.Flags().String("cacheDir", "", "filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/")
	modCmd.Flags().Bool("cleanDestinationDir", false, "remove files from destination not found in static directories")
	modCmd.Flags().StringP("contentDir", "c", "", "filesystem path to content directory")
	modCmd.Flags().StringP("destination", "d", "", "filesystem path to write files to")
	modCmd.Flags().StringSlice("disableKinds", []string{}, "disable different kind of pages (home, RSS etc.)")
	modCmd.Flags().Bool("enableGitInfo", false, "add Git revision, date and author info to the pages")
	modCmd.Flags().Bool("forceSyncStatic", false, "copy all files when static is changed.")
	modCmd.Flags().Bool("gc", false, "enable to run some cleanup tasks (remove unused cache files) after the build")
	modCmd.Flags().Bool("i18n-warnings", false, "print missing translations")
	modCmd.Flags().Bool("ignoreCache", false, "ignores the cache directory")
	modCmd.Flags().StringP("layoutDir", "l", "", "filesystem path to layout directory")
	modCmd.Flags().Bool("minify", false, "minify any supported output format (HTML, XML etc.)")
	modCmd.Flags().Bool("noChmod", false, "don't sync permission mode of files")
	modCmd.Flags().Bool("noTimes", false, "don't sync modification time of files")
	modCmd.Flags().Bool("path-warnings", false, "print warnings on duplicate target paths etc.")
	modCmd.Flags().String("poll", "", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes")
	modCmd.Flags().Bool("print-mem", false, "print memory usage to screen at intervals")
	modCmd.Flags().String("profile-cpu", "", "write cpu profile to `file`")
	modCmd.Flags().String("profile-mem", "", "write memory profile to `file`")
	modCmd.Flags().String("profile-mutex", "", "write Mutex profile to `file`")
	modCmd.Flags().Bool("templateMetrics", false, "display metrics about template executions")
	modCmd.Flags().Bool("templateMetricsHints", false, "calculate some improvement hints when combined with --templateMetrics")
	modCmd.Flags().StringSliceP("theme", "t", []string{}, "themes to use (located in /themes/THEMENAME/)")
	modCmd.Flags().String("trace", "", "write trace to `file` (not useful in general)")
	rootCmd.AddCommand(modCmd)

	carapace.Gen(modCmd).FlagCompletion(carapace.ActionMap{
		"cacheDir":      carapace.ActionDirectories(),
		"contentDir":    carapace.ActionDirectories(),
		"destination":   carapace.ActionDirectories(),
		"disableKinds":  action.ActionKinds(),
		"layoutDir":     carapace.ActionDirectories(),
		"profile-cpu":   carapace.ActionFiles(),
		"profile-mem":   carapace.ActionFiles(),
		"profile-mutex": carapace.ActionFiles(),
		"theme":         action.ActionThemes(modCmd),
		"trace":         carapace.ActionFiles(),
	})
}
