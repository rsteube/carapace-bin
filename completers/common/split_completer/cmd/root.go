package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "split",
	Short: "split a file into pieces",
	Long:  "https://linux.die.net/man/1/split",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("additional-suffix", "", "append an additional SUFFIX to file names")
	rootCmd.Flags().StringP("bytes", "b", "", "put SIZE bytes per output file")
	rootCmd.Flags().BoolS("d", "d", false, "use numeric suffixes starting at 0, not alphabetic")
	rootCmd.Flags().BoolP("elide-empty-files", "e", false, "do not generate empty output files with '-n'")
	rootCmd.Flags().String("filter", "", "write to shell COMMAND; file name is $FILE")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().String("hex-suffixes", "", "same as -x, but allow setting the start value")
	rootCmd.Flags().StringP("line-bytes", "C", "", "put at most SIZE bytes of records per output file")
	rootCmd.Flags().StringP("lines", "l", "", "put NUMBER lines/records per output file")
	rootCmd.Flags().StringP("number", "n", "", "generate CHUNKS output files; see explanation below")
	rootCmd.Flags().String("numeric-suffixes", "", "same as -d, but allow setting the start value")
	rootCmd.Flags().StringP("separator", "t", "", "use SEP instead of newline as the record separator;")
	rootCmd.Flags().StringP("suffix-length", "a", "", "generate suffixes of length N (default 2)")
	rootCmd.Flags().BoolP("unbuffered", "u", false, "immediately copy input to output with '-n r/...'")
	rootCmd.Flags().Bool("verbose", false, "print a diagnostic just before each")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolS("x", "x", false, "use hex suffixes starting at 0, not alphabetic")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
