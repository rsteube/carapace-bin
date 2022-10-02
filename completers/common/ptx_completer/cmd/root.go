package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ptx",
	Short: "produce a permuted index of file contents",
	Long:  "https://linux.die.net/man/1/ptx",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("O", "O", "", "generate output as roff directives")
	rootCmd.Flags().StringS("T", "T", "", "generate output as TeX directives")
	rootCmd.Flags().BoolP("auto-reference", "A", false, "output automatically generated references")
	rootCmd.Flags().StringP("break-file", "b", "", "word break characters in this FILE")
	rootCmd.Flags().StringP("flag-truncation", "F", "", "use STRING for flagging line truncations.")
	rootCmd.Flags().String("format", "", "generate output as given directives")
	rootCmd.Flags().StringP("gap-size", "g", "", "gap size in columns between output fields")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-case", "f", false, "fold lower case to upper case for sorting")
	rootCmd.Flags().StringP("ignore-file", "i", "", "read ignore word list from FILE")
	rootCmd.Flags().StringP("macro-name", "M", "", "macro name to use instead of 'xx'")
	rootCmd.Flags().StringP("only-file", "o", "", "read only word list from this FILE")
	rootCmd.Flags().BoolP("references", "r", false, "first field of each line is a reference")
	rootCmd.Flags().BoolP("right-side-refs", "R", false, "put references at right, not counted in -w")
	rootCmd.Flags().StringP("sentence-regexp", "S", "", "for end of lines or end of sentences")
	rootCmd.Flags().BoolP("traditional", "G", false, "behave more like System V 'ptx'")
	rootCmd.Flags().BoolP("typeset-mode", "t", false, "- not implemented -")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().StringP("width", "w", "", "output width in columns, reference excluded")
	rootCmd.Flags().StringP("word-regexp", "W", "", "use REGEXP to match each keyword")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"break-file":  carapace.ActionFiles(),
		"format":      carapace.ActionValues("roff", "tex"),
		"ignore-file": carapace.ActionFiles(),
		"only-file":   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("traditional").Changed {
				return carapace.ActionMessage("traditional mode only allows [INPUT [OUTPUT]]]")
			} else {
				return carapace.ActionFiles()
			}
		}),
	)
}
