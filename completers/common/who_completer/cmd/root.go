package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "who",
	Short: "show who is logged on",
	Long:  "https://linux.die.net/man/1/who",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "same as -b -d --login -p -r -t -T -u")
	rootCmd.Flags().BoolP("boot", "b", false, "time of last system boot")
	rootCmd.Flags().BoolP("count", "q", false, "all login names and number of users logged on")
	rootCmd.Flags().BoolP("dead", "d", false, "print dead processes")
	rootCmd.Flags().BoolP("heading", "H", false, "print line of column headings")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("login", "l", false, "print system login processes")
	rootCmd.Flags().Bool("lookup", false, "attempt to canonicalize hostnames via DNS")
	rootCmd.Flags().BoolS("m", "m", false, "only hostname and user associated with stdin")
	rootCmd.Flags().BoolP("mesg", "T", false, "add user's message status as +, - or ?")
	rootCmd.Flags().Bool("message", false, "same as -T")
	rootCmd.Flags().BoolP("process", "p", false, "print active processes spawned by init")
	rootCmd.Flags().BoolP("runlevel", "r", false, "print current runlevel")
	rootCmd.Flags().BoolP("short", "s", false, "print only name, line, and time (default)")
	rootCmd.Flags().BoolP("time", "t", false, "print last system clock change")
	rootCmd.Flags().BoolP("users", "u", false, "list users logged in")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("writable", "w", false, "same as -T")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
