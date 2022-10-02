package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/hugo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new content for your site",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	newCmd.Flags().StringP("baseURL", "b", "", "hostname (and path) to the root, e.g. http://spf13.com/")
	newCmd.Flags().BoolP("buildDrafts", "D", false, "include content marked as draft")
	newCmd.Flags().BoolP("buildExpired", "E", false, "include expired content")
	newCmd.Flags().BoolP("buildFuture", "F", false, "include content with publishdate in the future")
	newCmd.Flags().String("cacheDir", "", "filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/")
	newCmd.Flags().Bool("cleanDestinationDir", false, "remove files from destination not found in static directories")
	newCmd.Flags().StringP("contentDir", "c", "", "filesystem path to content directory")
	newCmd.Flags().StringP("destination", "d", "", "filesystem path to write files to")
	newCmd.Flags().StringSlice("disableKinds", []string{}, "disable different kind of pages (home, RSS etc.)")
	newCmd.Flags().String("editor", "", "edit new content with this editor, if provided")
	newCmd.Flags().Bool("enableGitInfo", false, "add Git revision, date and author info to the pages")
	newCmd.Flags().Bool("forceSyncStatic", false, "copy all files when static is changed.")
	newCmd.Flags().Bool("gc", false, "enable to run some cleanup tasks (remove unused cache files) after the build")
	newCmd.Flags().Bool("i18n-warnings", false, "print missing translations")
	newCmd.Flags().Bool("ignoreCache", false, "ignores the cache directory")
	newCmd.Flags().StringP("kind", "k", "", "content type to create")
	newCmd.Flags().StringP("layoutDir", "l", "", "filesystem path to layout directory")
	newCmd.Flags().Bool("minify", false, "minify any supported output format (HTML, XML etc.)")
	newCmd.Flags().Bool("noChmod", false, "don't sync permission mode of files")
	newCmd.Flags().Bool("noTimes", false, "don't sync modification time of files")
	newCmd.Flags().Bool("path-warnings", false, "print warnings on duplicate target paths etc.")
	newCmd.Flags().String("poll", "", "set this to a poll interval, e.g --poll 700ms, to use a poll based approach to watch for file system changes")
	newCmd.Flags().Bool("print-mem", false, "print memory usage to screen at intervals")
	newCmd.Flags().String("profile-cpu", "", "write cpu profile to `file`")
	newCmd.Flags().String("profile-mem", "", "write memory profile to `file`")
	newCmd.Flags().String("profile-mutex", "", "write Mutex profile to `file`")
	newCmd.Flags().Bool("templateMetrics", false, "display metrics about template executions")
	newCmd.Flags().Bool("templateMetricsHints", false, "calculate some improvement hints when combined with --templateMetrics")
	newCmd.Flags().StringSliceP("theme", "t", []string{}, "themes to use (located in /themes/THEMENAME/)")
	newCmd.Flags().String("trace", "", "write trace to `file` (not useful in general)")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"cacheDir":     carapace.ActionDirectories(),
		"contentDir":   carapace.ActionDirectories(),
		"destination":  carapace.ActionDirectories(),
		"disableKinds": action.ActionKinds(),
		"editor":       carapace.ActionFiles(),
		"layoutDir":    carapace.ActionDirectories(),
		"theme":        action.ActionThemes(newCmd),
		"trace":        carapace.ActionFiles(),
	})
}
