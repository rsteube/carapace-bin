package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "more",
	Short: "file perusal filter for crt viewing",
	Long:  "https://linux.die.net/man/1/more",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("clean-print", "p", false, "do not scroll, clean screen and display text")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("lines", "n", "", "the number of lines per screenful")
	rootCmd.Flags().BoolP("logical", "f", false, "count logical rather than screen lines")
	rootCmd.Flags().BoolP("no-pause", "l", false, "suppress pause after form feed")
	rootCmd.Flags().BoolP("plain", "u", false, "suppress underlining and bold")
	rootCmd.Flags().BoolP("print-over", "c", false, "do not scroll, display text and clean line ends")
	rootCmd.Flags().BoolP("silent", "d", false, "display help instead of ringing bell")
	rootCmd.Flags().BoolP("squeeze", "s", false, "squeeze multiple blank lines into one")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
