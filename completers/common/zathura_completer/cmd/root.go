package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zathura",
	Short: "a document viewer",
	Long:  "https://pwmt.org/projects/zathura/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bookmark", "b", "", "Bookmark to go to")
	rootCmd.Flags().String("cache-dir", "", "Path to the cache directory")
	rootCmd.Flags().StringP("config-dir", "c", "", "Path to the config directory")
	rootCmd.Flags().StringP("data-dir", "d", "", "Path to the data directory")
	rootCmd.Flags().StringP("find", "f", "", "Search for the given phrase and display results")
	rootCmd.Flags().Bool("fork", false, "Fork into the background")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().StringP("log-level", "l", "", "Log level (debug, info, warning, error)")
	rootCmd.Flags().String("mode", "", "Start in a non-default mode")
	rootCmd.Flags().StringP("page", "P", "", "Page number to go to")
	rootCmd.Flags().StringP("password", "w", "", "Document password")
	rootCmd.Flags().StringP("plugins-dir", "p", "", "Path to the directories containing plugins")
	rootCmd.Flags().StringP("reparent", "e", "", "Reparents to window specified by xid (X11)")
	rootCmd.Flags().StringP("synctex-editor-command", "x", "", "Synctex editor (forwarded to the synctex command)")
	rootCmd.Flags().String("synctex-forward", "", "Move to given synctex position")
	rootCmd.Flags().String("synctex-pid", "", "Highlight given position in the given process")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config-dir":  carapace.ActionDirectories(),
		"data-dir":    carapace.ActionDirectories(),
		"log-level":   carapace.ActionValues("debug", "info", "warning", "error").StyleF(style.ForLogLevel),
		"mode":        carapace.ActionValues("fullscreen", "presentation"),
		"plugins-dir": carapace.ActionDirectories(),
		"synctex-pid": ps.ActionProcessIds(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
