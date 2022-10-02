package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lazygit",
	Short: "simple terminal UI for git commands",
	Long:  "https://github.com/jesseduffield/lazygit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("config", "c", false, "Print the default config")
	rootCmd.Flags().BoolP("debug", "d", false, "Run in debug mode with logging (see --logs flag below). Use the LOG_LEVEL env var to set the log level (debug/info/warn/error)")
	rootCmd.Flags().StringP("filter", "f", "", "Path to filter on in `git log -- <path>`.")
	rootCmd.Flags().StringP("git-dir", "g", "", "equivalent of the --git-dir git argument")
	rootCmd.Flags().BoolP("help", "h", false, "Displays help with available flag, subcommand, and positional value parameters.")
	rootCmd.Flags().BoolP("logs", "l", false, "Tail lazygit logs (intended to be used when `lazygit --debug` is called in a separate terminal tab)")
	rootCmd.Flags().StringP("path", "p", "", "Path of git repo.")
	rootCmd.Flags().Bool("print-config-dir", false, "Print the config directory")
	rootCmd.Flags().String("use-config-dir", "", "override default config directory with provided directory")
	rootCmd.Flags().String("use-config-file", "", "Comma seperated list to custom config file(s)")
	rootCmd.Flags().BoolP("version", "v", false, "Print the current version")
	rootCmd.Flags().StringP("work-tree", "w", "", "equivalent of the --work-tree git argument")

	// TODO work-tree, filter
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"git-dir":        carapace.ActionDirectories(),
		"path":           carapace.ActionDirectories(),
		"use-config-dir": carapace.ActionDirectories(),
		"use-config-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	})
}
