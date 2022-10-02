package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tr",
	Short: "translate or delete characters",
	Long:  "https://linux.die.net/man/1/tr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("delete", "d", false, "delete characters in SET1, do not translate")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("squeeze-repeats", "s", false, "replace each sequence of a repeated character")
	rootCmd.Flags().BoolP("truncate-set1", "t", false, "first truncate SET1 to length of SET2")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
}
