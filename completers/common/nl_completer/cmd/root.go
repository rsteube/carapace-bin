package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nl",
	Short: "number lines of files",
	Long:  "https://linux.die.net/man/1/nl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("body-numbering", "b", "", "use STYLE for numbering body lines")
	rootCmd.Flags().StringP("footer-numbering", "f", "", "use STYLE for numbering footer lines")
	rootCmd.Flags().StringP("header-numbering", "h", "", "use STYLE for numbering header lines")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("join-blank-lines", "l", "", "group of NUMBER empty lines counted as one")
	rootCmd.Flags().StringP("line-increment", "i", "", "line number increment at each line")
	rootCmd.Flags().BoolP("no-renumber", "p", false, "do not reset line numbers for each section")
	rootCmd.Flags().StringP("number-format", "n", "", "insert line numbers according to FORMAT")
	rootCmd.Flags().StringP("number-separator", "s", "", "add STRING after (possible) line number")
	rootCmd.Flags().StringP("number-width", "w", "", "use NUMBER columns for line numbers")
	rootCmd.Flags().StringP("section-delimiter", "d", "", "use CC for logical page delimiters")
	rootCmd.Flags().StringP("starting-line-number", "v", "", "first line number for each section")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"body-numbering":   ActionNumbering(),
		"footer-numbering": ActionNumbering(),
		"header-numbering": ActionNumbering(),
		"number-format": carapace.ActionValuesDescribed(
			"ln", "left justified",
			"rn", "right justified",
			"rz", "right justified with leading zeroes",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}

func ActionNumbering() carapace.Action {
	return carapace.ActionValuesDescribed(
		"a", "number all lines",
		"n", "no line numbering",
		"p", "number only lines matching regex",
		"t", "number only non-empty lines",
	)
}
