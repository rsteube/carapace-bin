package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kak"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kak",
	Short: "a vim-inspired, selection oriented code editor",
	Long:  "http://kakoune.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("E", "E", "", "execute argument on server initialisation")
	rootCmd.Flags().StringS("c", "c", "", "connect to given session")
	rootCmd.Flags().BoolS("clear", "clear", false, "clear dead sessions")
	rootCmd.Flags().BoolS("d", "d", false, "run as a headless session (requires -s)")
	rootCmd.Flags().StringS("e", "e", "", "execute argument on client initialisation")
	rootCmd.Flags().StringS("f", "f", "", "filter: for each file, select the entire buffer and execute the given keys")
	rootCmd.Flags().BoolS("help", "help", false, "display a help message and quit")
	rootCmd.Flags().StringS("i", "i", "", "backup the files on which a filter is applied using the given suffix")
	rootCmd.Flags().BoolS("l", "l", false, "list existing sessions")
	rootCmd.Flags().BoolS("n", "n", false, "do not source kakrc files on startup")
	rootCmd.Flags().StringS("p", "p", "", "just send stdin as commands to the given session")
	rootCmd.Flags().BoolS("q", "q", false, "in filter mode, be quiet about errors applying keys")
	rootCmd.Flags().BoolS("ro", "ro", false, "readonly mode")
	rootCmd.Flags().StringS("s", "s", "", "set session name")
	rootCmd.Flags().StringS("ui", "ui", "", "set the type of user interface to use (terminal, dummy, or json)")
	rootCmd.Flags().BoolS("version", "version", false, "display kakoune version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c":  kak.ActionSessions(),
		"p":  kak.ActionSessions(),
		"ui": carapace.ActionValues("terminal", "dummy", "json"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
