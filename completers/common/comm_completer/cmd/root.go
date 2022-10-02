package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "comm",
	Short: "compare two sorted files line by line",
	Long:  "https://linux.die.net/man/1/comm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "suppress column 1 (lines unique to FILE1)")
	rootCmd.Flags().BoolS("2", "2", false, "suppress column 2 (lines unique to FILE2)")
	rootCmd.Flags().BoolS("3", "3", false, "suppress column 3 (lines that appear in both files)")
	rootCmd.Flags().Bool("check-order", false, "check that the input is correctly sorted")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("nocheck-order", false, "do not check that the input is correctly sorted")
	rootCmd.Flags().String("output-delimiter", "", "separate columns with STR")
	rootCmd.Flags().Bool("total", false, "output a summary")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
