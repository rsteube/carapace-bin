package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "concatenate files and print on the standard output",
	Long:  "https://linux.die.net/man/1/cat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "equivalent to -vE")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("number", "n", false, "number all output lines")
	rootCmd.Flags().BoolP("number-nonblank", "b", false, "number nonempty output lines, overrides -n")
	rootCmd.Flags().BoolP("show-all", "A", false, "equivalent to -vET")
	rootCmd.Flags().BoolP("show-ends", "E", false, "display $ at end of each line")
	rootCmd.Flags().BoolP("show-nonprinting", "v", false, "use ^ and M- notation, except for LFD and TAB")
	rootCmd.Flags().BoolP("show-tabs", "T", false, "display TAB characters as ^I")
	rootCmd.Flags().BoolP("squeeze-blank", "s", false, "suppress repeated empty output lines")
	rootCmd.Flags().BoolS("t", "t", false, "equivalent to -vT")
	rootCmd.Flags().BoolS("u", "u", false, "(ignored)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
