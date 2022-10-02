package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "watch",
	Short: "execute a program periodically, showing output fullscreen",
	Long:  "https://linux.die.net/man/1/watch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("beep", "b", false, "beep if command has a non-zero exit")
	rootCmd.Flags().BoolP("chgexit", "g", false, "exit when output from command changes")
	rootCmd.Flags().BoolP("color", "c", false, "interpret ANSI color and style sequences")
	rootCmd.Flags().StringP("differences", "d", "", "highlight changes between updates")
	rootCmd.Flags().BoolP("errexit", "e", false, "exit if command has a non-zero exit")
	rootCmd.Flags().BoolP("exec", "x", false, "pass command to exec instead of \"sh -c\"")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().StringP("interval", "n", "", "seconds to wait between updates")
	rootCmd.Flags().BoolP("no-title", "t", false, "turn off header")
	rootCmd.Flags().BoolP("no-wrap", "w", false, "turn off line wrapping")
	rootCmd.Flags().BoolP("precise", "p", false, "attempt run command in precise intervals")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")

	rootCmd.Flag("differences").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"differences": carapace.ActionValues("permanent"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
