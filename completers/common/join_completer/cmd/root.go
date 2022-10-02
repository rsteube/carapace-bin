package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "join",
	Short: "join lines of two files on a common field",
	Long:  "https://linux.die.net/man/1/join",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("1", "1", "", "join on this FIELD of file 1")
	rootCmd.Flags().StringS("2", "2", "", "join on this FIELD of file 2")
	rootCmd.Flags().StringS("a", "a", "", "also print unpairable lines from file FILENUM, where")
	rootCmd.Flags().Bool("check-order", false, "check that the input is correctly sorted, even")
	rootCmd.Flags().StringS("e", "e", "", "replace missing input fields with EMPTY")
	rootCmd.Flags().Bool("header", false, "treat the first line in each file as field headers,")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "ignore differences in case when comparing fields")
	rootCmd.Flags().StringS("j", "j", "", "equivalent to '-1 FIELD -2 FIELD'")
	rootCmd.Flags().Bool("nocheck-order", false, "do not check that the input is correctly sorted")
	rootCmd.Flags().StringS("o", "o", "", "obey FORMAT while constructing output line")
	rootCmd.Flags().StringS("t", "t", "", "use CHAR as input and output field separator")
	rootCmd.Flags().StringS("v", "v", "", "like -a FILENUM, but suppress joined output lines")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"a": carapace.ActionValues("1", "2"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
