package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "strings",
	Short: "print the sequences of printable characters in files",
	Long:  "https://linux.die.net/man/1/strings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Scan the entire file, not just the data section [default]")
	rootCmd.Flags().StringP("bytes", "n", "", "Locate & print any NUL-terminated sequence of at")
	rootCmd.Flags().BoolP("data", "d", false, "Only scan the data sections in the file")
	rootCmd.Flags().StringP("encoding", "e", "", "Select character size and endianness:")
	rootCmd.Flags().BoolP("help", "h", false, "Display this information")
	rootCmd.Flags().BoolP("include-all-whitespace", "w", false, "Include all whitespace as valid string characters")
	rootCmd.Flags().BoolS("o", "o", false, "An alias for --radix=o")
	rootCmd.Flags().StringP("output-separator", "s", "", "String used to separate strings in output.")
	rootCmd.Flags().BoolP("print-file-name", "f", false, "Print the name of the file before each string")
	rootCmd.Flags().StringP("radix", "t", "", "Print the location of the string in base 8, 10 or 16")
	rootCmd.Flags().StringP("target", "T", "", "Specify the binary file format")
	rootCmd.Flags().BoolP("version", "v", false, "Print the program's version number")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"encoding": carapace.ActionValuesDescribed(
			"s", "7-bit",
			"S", "8-bit",
			"b", "16-bit",
			"l", "16-bit",
			"B", "32-bit",
			"L", "32-bit",
		),
		"radix": carapace.ActionValuesDescribed(
			"o", "8",
			"d", "10",
			"x", "16",
		),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
