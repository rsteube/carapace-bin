package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ranger",
	Short: "visual file manager",
	Long:  "https://github.com/ranger/ranger",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cachedir", "", "change the cache directory.")
	rootCmd.Flags().String("choosedir", "", "Makes ranger act like a directory chooser.")
	rootCmd.Flags().String("choosefile", "", "Makes ranger act like a file chooser.")
	rootCmd.Flags().String("choosefiles", "", "Makes ranger act like a file chooser for multiple files at once.")
	rootCmd.Flags().BoolP("clean", "c", false, "don't touch/require any config files.")
	rootCmd.Flags().String("cmd", "", "Execute COMMAND after the configuration has been read.")
	rootCmd.Flags().String("copy-config", "", "copy the default configs to the local config directory.")
	rootCmd.Flags().String("datadir", "", "change the data directory.")
	rootCmd.Flags().BoolP("debug", "d", false, "activate debug mode")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().String("list-tagged-files", "", "List all files which are tagged with the given tag.")
	rootCmd.Flags().Bool("list-unused-keys", false, "List common keys which are not bound to any action.")
	rootCmd.Flags().String("logfile", "", "log file to use, '-' for stderr")
	rootCmd.Flags().Bool("profile", false, "Print statistics of CPU usage on exit.")
	rootCmd.Flags().String("selectfile", "", "Open ranger with supplied file selected.")
	rootCmd.Flags().Bool("show-only-dirs", false, "Show only directories, no files or links")
	rootCmd.Flags().Bool("version", false, "show program's version number and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cachedir":    carapace.ActionDirectories(),
		"choosedir":   carapace.ActionFiles(),
		"choosefile":  carapace.ActionFiles(),
		"choosefiles": carapace.ActionFiles(),
		"copy-config": carapace.ActionValues("all", "rc", "rifle", "commands", "commands_full", "scope"),
		"datadir":     carapace.ActionDirectories(),
		"logfile":     carapace.ActionFiles(),
		"selectfile":  carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(carapace.ActionFiles())
}
