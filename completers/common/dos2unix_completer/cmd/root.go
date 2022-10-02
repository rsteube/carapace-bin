package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dos2unix",
	Short: "DOS/Mac to Unix and vice versa text file format converter",
	Long:  "https://en.wikipedia.org/wiki/Unix2dos",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("7", "7", false, "convert 8 bit characters to 7 bit space")
	rootCmd.Flags().BoolP("add-bom", "m", false, "add Byte Order Mark (default UTF-8)")
	rootCmd.Flags().Bool("allow-chown", false, "allow file ownership change")
	rootCmd.Flags().StringP("convmode", "c", "", "conversion mode")
	rootCmd.Flags().StringP("display-enc", "D", "", "Set encoding of displayed text.")
	rootCmd.Flags().BoolP("follow-symlink", "F", false, "follow symbolic links and convert the targets")
	rootCmd.Flags().BoolP("force", "f", false, "force conversion of binary files")
	rootCmd.Flags().BoolP("help", "h", false, "display this help text")
	rootCmd.Flags().StringP("info", "i", "", "display file information")
	rootCmd.Flags().BoolP("keep-bom", "b", false, "keep Byte Order Mark")
	rootCmd.Flags().BoolP("keep-utf16", "u", false, "keep UTF-16 encoding")
	rootCmd.Flags().BoolP("keepdate", "k", false, "keep output file date")
	rootCmd.Flags().BoolP("license", "L", false, "display software license")
	rootCmd.Flags().BoolP("newfile", "n", false, "write to new file")
	rootCmd.Flags().BoolP("newline", "l", false, "add additional newline")
	rootCmd.Flags().Bool("no-allow-chown", false, "don't allow file ownership change (default)")
	rootCmd.Flags().BoolP("oldfile", "o", false, "write to old file (default)")
	rootCmd.Flags().BoolP("quiet", "q", false, "quiet mode, suppress all warnings")
	rootCmd.Flags().BoolP("remove-bom", "r", false, "remove Byte Order Mark (default)")
	rootCmd.Flags().BoolP("safe", "s", false, "skip binary files (default)")
	rootCmd.Flags().BoolP("skip-symlink", "S", false, "keep symbolic links and targets unchanged (default)")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose operation")
	rootCmd.Flags().BoolP("version", "V", false, "display version number")

	rootCmd.Flag("info").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"convmode":    carapace.ActionValues("ascii", "7bit", "iso", "mac"),
		"display-enc": carapace.ActionValues("ansi", "unicode", "unicodebom", "utf8", "utf8bom"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
