package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "print newline, word, and byte counts for each file",
	Long:  "https://linux.die.net/man/1/wc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("bytes", "c", false, "print the byte counts")
	rootCmd.Flags().BoolP("chars", "m", false, "print the character counts")
	rootCmd.Flags().String("files0-from", "", "read input from the files specified by")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("lines", "l", false, "print the newline counts")
	rootCmd.Flags().BoolP("max-line-length", "L", false, "print the maximum display width")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("words", "w", false, "print the word counts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"files0-from": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("files0-from").Changed {
				return carapace.ActionMessage("files0-from used")
			} else {
				return carapace.ActionFiles()
			}
		}),
	)
}
