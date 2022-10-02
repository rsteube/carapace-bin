package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kak"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kak-lsp",
	Short: "Kakoune Language Server Protocol Client",
	Long:  "https://github.com/kak-lsp/kak-lsp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "Read config from FILE")
	rootCmd.Flags().BoolP("daemonize", "d", false, "Daemonize kak-lsp process (server only)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().Bool("initial-request", false, "Read initial request from stdin")
	rootCmd.Flags().Bool("kakoune", false, "Generate commands for Kakoune to plug in kak-lsp")
	rootCmd.Flags().String("log", "", "File to write the log into instead of stderr")
	rootCmd.Flags().Bool("request", false, "Forward stdin to kak-lsp server")
	rootCmd.Flags().StringP("session", "s", "", "Session id to communicate via unix socket")
	rootCmd.Flags().StringP("timeout", "t", "", "Session timeout in seconds (default 1800)")
	rootCmd.Flags().BoolS("v", "v", false, "Sets the level of verbosity (use up to 4 times)")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config":  carapace.ActionFiles(),
		"log":     carapace.ActionFiles(),
		"session": kak.ActionSessions(),
	})
}
