package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "diff3",
	Short: "compare three files line by line",
	Long:  "https://linux.die.net/man/1/diff3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("X", "X", false, "like -x, but bracket conflicts")
	rootCmd.Flags().String("diff-program", "", "use PROGRAM to compare files")
	rootCmd.Flags().BoolP("easy-only", "3", false, "like -e, but incorporate only nonoverlapping changes")
	rootCmd.Flags().BoolP("ed", "e", false, "output ed script incorporating changes")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolS("i", "i", false, "append 'w' and 'q' commands to ed scripts")
	rootCmd.Flags().BoolP("initial-tab", "T", false, "make tabs line up by prepending a tab")
	rootCmd.Flags().StringArrayP("label", "L", []string{}, "use LABEL instead of file name")
	rootCmd.Flags().BoolP("merge", "m", false, "output actual merged file")
	rootCmd.Flags().BoolP("overlap-only", "x", false, "like -e, but incorporate only overlapping changes")
	rootCmd.Flags().BoolP("show-all", "A", false, "output all changes, bracketing conflicts")
	rootCmd.Flags().BoolP("show-overlap", "E", false, "like -e, but bracket conflicts")
	rootCmd.Flags().Bool("strip-trailing-cr", false, "strip trailing carriage return on input")
	rootCmd.Flags().BoolP("text", "a", false, "treat all files as text")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"diff-program": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
