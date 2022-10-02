package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pr",
	Short: "convert text files for printing",
	Long:  "https://linux.die.net/man/1/pr",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("F", "F", "", "use form feeds instead of newlines to separate pages")
	rootCmd.Flags().BoolP("across", "a", false, "print columns across rather than down, used together")
	rootCmd.Flags().String("columns", "", "output COLUMN columns and print columns down")
	rootCmd.Flags().StringP("date-format", "D", "", "use FORMAT for the header date")
	rootCmd.Flags().BoolP("double-space", "d", false, "double space the output")
	rootCmd.Flags().StringP("expand", "e", "", "expand input CHARs (TABs) to tab WIDTH (8)")
	rootCmd.Flags().StringS("first-line-number", "N", "", "start counting with NUMBER at 1st line of first page printed")
	rootCmd.Flags().StringP("form-feed", "f", "", "use form feeds instead of newlines to separate pages")
	rootCmd.Flags().StringP("header", "h", "", "use a centered HEADER instead of filename in page header")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("indent", "o", "", "offset each line with MARGIN (zero) spaces,")
	rootCmd.Flags().BoolP("join-lines", "J", false, "merge full lines, turns off -W line truncation, no column")
	rootCmd.Flags().StringP("length", "l", "", "set the page length to PAGE_LENGTH (66) lines")
	rootCmd.Flags().BoolP("merge", "m", false, "print all files in parallel, one in each column,")
	rootCmd.Flags().BoolP("no-file-warnings", "r", false, "omit warning when a file cannot be opened")
	rootCmd.Flags().StringP("number-lines", "n", "", "number lines, use DIGITS (5) digits")
	rootCmd.Flags().BoolP("omit-header", "t", false, "omit page headers and trailers;")
	rootCmd.Flags().BoolP("omit-pagination", "T", false, "omit page headers and trailers, eliminate any pagination")
	rootCmd.Flags().StringP("output-tabs", "i", "", "replace spaces with CHARs (TABs) to tab WIDTH (8)")
	rootCmd.Flags().StringP("page-width", "W", "", "set page width to PAGE_WIDTH (72) characters alway")
	rootCmd.Flags().String("pages", "", "begin [stop] printing with page")
	rootCmd.Flags().StringP("sep-string", "S", "", "separate columns by STRING")
	rootCmd.Flags().StringP("separator", "s", "", "separate columns by a single character")
	rootCmd.Flags().BoolP("show-control-chars", "c", false, "use hat notation (^G) and octal backslash notation")
	rootCmd.Flags().BoolP("show-nonprinting", "v", false, "use octal backslash notation")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "w", "", "set page width to PAGE_WIDTH (72) characters")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
