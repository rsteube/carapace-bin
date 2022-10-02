package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uniq",
	Short: "report or omit repeated lines",
	Long:  "https://linux.die.net/man/1/uniq",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "print all duplicate lines")
	rootCmd.Flags().String("all-repeated", "", "like -D, but allow separating groups")
	rootCmd.Flags().StringP("check-chars", "w", "", "compare no more than N characters in lines")
	rootCmd.Flags().BoolP("count", "c", false, "prefix lines by the number of occurrences")
	rootCmd.Flags().String("group", "", "show all items, separating groups with an empty line;")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "ignore differences in case when comparing")
	rootCmd.Flags().BoolP("repeated", "d", false, "only print duplicate lines, one for each group")
	rootCmd.Flags().StringP("skip-chars", "s", "", "avoid comparing the first N characters")
	rootCmd.Flags().StringP("skip-fields", "f", "", "avoid comparing the first N fields")
	rootCmd.Flags().BoolP("unique", "u", false, "only print unique lines")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"all-repeated": carapace.ActionValues("none", "prepend", "separate"),
		"group":        carapace.ActionValues("append", "both", "prepend", "separate"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
