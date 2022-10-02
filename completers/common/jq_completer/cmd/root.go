package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jq",
	Short: "Command-line JSON processor",
	Long:  "https://stedolan.github.io/jq/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("L", "L", "", "prepend a directory to the module search path")
	rootCmd.Flags().Bool("arg", false, "pre-set a variable to a string")
	rootCmd.Flags().Bool("argjson", false, "pre-set a variable to an object")
	rootCmd.Flags().Bool("args", false, "remaining arguments are string arguments, not files")
	rootCmd.Flags().BoolP("ascii-output", "a", false, "restrict output to ASCII")
	rootCmd.Flags().BoolP("color-output", "C", false, "output in color")
	rootCmd.Flags().BoolP("compact-output", "c", false, "don't pretty-print")
	rootCmd.Flags().BoolP("exit-status", "e", false, "report \"false\" and \"null\" results via exit code")
	rootCmd.Flags().StringP("from-file", "f", "", "read filter from file")
	rootCmd.Flags().String("indent", "", "indent output using given number of spaces")
	rootCmd.Flags().BoolP("join-output", "j", false, "like -r, without newlines between outputs")
	rootCmd.Flags().Bool("jsonargs", false, "remaining arguments are JSON arguments, not files")
	rootCmd.Flags().BoolP("monochrome-output", "M", false, "output without color")
	rootCmd.Flags().BoolP("null-input", "n", false, "input is ignored")
	rootCmd.Flags().BoolP("raw-input", "R", false, "consider each input line as a JSON strings")
	rootCmd.Flags().BoolP("raw-output", "r", false, "don't JSON-quote output if it's a string")
	rootCmd.Flags().Bool("rawfile", false, "pre-set a variable to the contents of a file")
	rootCmd.Flags().Bool("seq", false, "use application/json-seq ASCII RS/LF scheme in input and output")
	rootCmd.Flags().BoolP("slurp", "s", false, "join input JSON objects to array before filtering")
	rootCmd.Flags().Bool("slurpfile", false, "pre-set a variable to an array of JSON texts read from a file")
	rootCmd.Flags().BoolP("sort-keys", "S", false, "output object keys in sorted order")
	rootCmd.Flags().Bool("stream", false, "parse input streamily (changes output)")
	rootCmd.Flags().Bool("tab", false, "indent output using TAB characters")
	rootCmd.Flags().Bool("unbuffered", false, "flush output after each JSON object")
	rootCmd.Flags().Bool("version", false, "output jq's version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"L":         carapace.ActionDirectories(),
		"from-file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("args").Changed || rootCmd.Flag("jsonargs").Changed {
				return carapace.ActionValues()
			} else {
				return carapace.ActionFiles()
			}
		}),
	)
}
