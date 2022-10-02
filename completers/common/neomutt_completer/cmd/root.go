package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "neomutt",
	Short: "The NeoMutt Mail User Agent",
	Long:  "https://neomutt.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "Print an expanded version of the given alias to stdout and exit")
	rootCmd.Flags().BoolS("B", "B", false, "Run in batch mode (do not start the ncurses UI)")
	rootCmd.Flags().BoolS("D", "D", false, "Dump all config variables as 'name=value' pairs to stdout")
	rootCmd.Flags().BoolS("E", "E", false, "Edit draft (-H) or include (-i) file during message composition")
	rootCmd.Flags().StringS("F", "F", "", "Specify an alternative initialization file to read")
	rootCmd.Flags().BoolS("G", "G", false, "Start NeoMutt with a listing of subscribed newsgroups")
	rootCmd.Flags().StringS("H", "H", "", "Specify a draft file with header and body for message composing")
	rootCmd.Flags().BoolS("R", "R", false, "Open mailbox in read-only mode")
	rootCmd.Flags().BoolS("Z", "Z", false, "Open the first mailbox with new message or exit immediately")
	rootCmd.Flags().StringS("a", "a", "", "Attach one or more files to a message (must be the last option)")
	rootCmd.Flags().StringS("b", "b", "", "Specify a blind carbon copy (Bcc) recipient")
	rootCmd.Flags().StringS("c", "c", "", "Specify a carbon copy (Cc) recipient")
	rootCmd.Flags().StringS("d", "d", "", "Log debugging output to a file (default is \"~/.neomuttdebug0\")")
	rootCmd.Flags().StringS("e", "e", "", "Specify a command to be run after reading the config files")
	rootCmd.Flags().StringS("f", "f", "", "Specify a mailbox (as defined with 'mailboxes' command) to load")
	rootCmd.Flags().StringS("g", "g", "", "Like -G, but start at specified news server")
	rootCmd.Flags().BoolS("h", "h", false, "Print this help message and exit")
	rootCmd.Flags().StringS("i", "i", "", "Specify an include file to be embedded in the body of a message")
	rootCmd.Flags().StringS("l", "l", "", "Specify a file for debugging output (default \"~/.neomuttdebug0\")")
	rootCmd.Flags().StringS("m", "m", "", "Specify a default mailbox format type for newly created folders")
	rootCmd.Flags().BoolS("n", "n", false, "Do not read the system-wide configuration file")
	rootCmd.Flags().BoolS("p", "p", false, "Resume a prior postponed message, if any")
	rootCmd.Flags().StringS("s", "s", "", "Specify a subject (must be enclosed in quotes if it has spaces)")
	rootCmd.Flags().BoolS("v", "v", false, "Print the NeoMutt version and compile-time definitions and exit")
	rootCmd.Flags().BoolS("y", "y", false, "Start NeoMutt with a listing of all defined mailboxes")
	rootCmd.Flags().BoolS("z", "z", false, "Open the first or specified (-f) mailbox if it holds any message")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionFiles(),
		"H": carapace.ActionFiles(),
		"a": carapace.ActionFiles(),
		"d": carapace.ActionValues("1", "2", "3", "4", "5"),
		"e": carapace.ActionFiles(),
		"i": carapace.ActionFiles(),
		"l": carapace.ActionFiles(),
		"m": carapace.ActionValues("MH", "MMDF", "Maildir", "mbox"),
	})
}
