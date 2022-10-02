package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paste",
	Short: "merge lines of files",
	Long:  "https://linux.die.net/man/1/paste",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delimiters", "d", "", "reuse characters from LIST instead of TABs")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("serial", "s", false, "paste one file at a time instead of in parallel")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
