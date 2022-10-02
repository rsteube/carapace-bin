package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pathchk",
	Short: "check whether file names are valid or portable",
	Long:  "https://linux.die.net/man/1/pathchk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "check for empty names and leading \"-\"")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("p", "p", false, "check for most POSIX systems")
	rootCmd.Flags().Bool("portability", false, "check for all POSIX systems (equivalent to -p -P)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
