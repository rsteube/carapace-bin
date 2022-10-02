package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "run remote shell command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellCmd).Standalone()

	shellCmd.Flags().BoolS("T", "T", false, "disable pty allocation")
	shellCmd.Flags().StringS("e", "e", "", "choose escape character, or \"none\"; default '~'")
	shellCmd.Flags().BoolS("n", "n", false, "don't read from stdin")
	shellCmd.Flags().BoolS("t", "t", false, "allocate a pty if on a tty (-tt: force pty allocation)")
	shellCmd.Flags().BoolS("x", "x", false, "disable remote exit codes and stdout/stderr separation")
	rootCmd.AddCommand(shellCmd)
}
