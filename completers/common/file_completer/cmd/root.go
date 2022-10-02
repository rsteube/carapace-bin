package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "file",
	Short: "determine file type",
	Long:  "https://linux.die.net/man/1/file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("apple", false, "output the Apple CREATOR/TYPE")
	rootCmd.Flags().BoolP("brief", "b", false, "do not prepend filenames to output lines")
	rootCmd.Flags().BoolP("checking-printout", "c", false, "print the parsed form of the magic file")
	rootCmd.Flags().BoolP("compile", "C", false, "compile file specified by -m")
	rootCmd.Flags().BoolP("debug", "d", false, "print debugging messages")
	rootCmd.Flags().BoolP("dereference", "L", false, "follow symlinks")
	rootCmd.Flags().StringP("exclude", "e", "", "exclude TEST from the list of test to be performed for file")
	rootCmd.Flags().String("exclude-quiet", "", "like exclude, but ignore unknown tests")
	rootCmd.Flags().Bool("extension", false, "output a slash-separated list of extensions")
	rootCmd.Flags().StringP("files-from", "f", "", "read the filenames to be examined from FILE")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("keep-going", "k", false, "don't stop at the first match")
	rootCmd.Flags().BoolP("list", "l", false, "list magic strength")
	rootCmd.Flags().StringP("magic-file", "m", "", "use LIST as a colon-separated list of magic number files")
	rootCmd.Flags().BoolP("mime", "i", false, "output MIME type strings")
	rootCmd.Flags().Bool("mime-encoding", false, "output the MIME encoding")
	rootCmd.Flags().Bool("mime-type", false, "output the MIME type")
	rootCmd.Flags().BoolP("no-buffer", "n", false, "do not buffer output")
	rootCmd.Flags().BoolP("no-dereference", "h", false, "don't follow symlinks (default)")
	rootCmd.Flags().BoolP("no-pad", "N", false, "do not pad output")
	rootCmd.Flags().BoolP("no-sandbox", "S", false, "disable system call sandboxing")
	rootCmd.Flags().BoolP("parameter", "P", false, "set file engine parameter limits")
	rootCmd.Flags().BoolP("preserve-date", "p", false, "preserve access times on files")
	rootCmd.Flags().BoolP("print0", "0", false, "terminate filenames with ASCII NUL")
	rootCmd.Flags().BoolP("raw", "r", false, "don't translate unprintable chars to \\ooo")
	rootCmd.Flags().StringP("separator", "F", "", "use string as separator instead of `:'")
	rootCmd.Flags().BoolP("special-files", "s", false, "treat special (block/char devices) files as ordinary ones")
	rootCmd.Flags().BoolP("uncompress", "z", false, "try to look inside compressed files")
	rootCmd.Flags().BoolP("uncompress-noreport", "Z", false, "only print the contents of compressed files")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exclude":       ActionTests(),
		"exclude-quiet": ActionTests(),
		"files-from":    carapace.ActionFiles(),
		"magic-file": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionFiles()
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

func ActionTests() carapace.Action {
	return carapace.ActionValues("apptype", "ascii", "cdf", "compress", "csv", "elf", "encoding", "soft", "tar", "json", "text", "tokens")
}
