package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "id",
	Short: "Print user and group information",
	Long:  "https://linux.die.net/man/1/id",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "ignore, for compatibility with other versions")
	rootCmd.Flags().BoolP("context", "Z", false, "print only the security context of the process")
	rootCmd.Flags().BoolP("group", "g", false, "print only the effective group ID")
	rootCmd.Flags().BoolP("groups", "G", false, "print all group IDs")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("name", "n", false, "print a name instead of a number, for -ugG")
	rootCmd.Flags().BoolP("real", "r", false, "print the real ID instead of the effective ID, with -ugG")
	rootCmd.Flags().BoolP("user", "u", false, "print only the effective user ID")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero", "z", false, "delimit entries with NUL characters, not whitespace;")

	carapace.Gen(rootCmd).PositionalCompletion(
		os.ActionUsers(),
	)
}
