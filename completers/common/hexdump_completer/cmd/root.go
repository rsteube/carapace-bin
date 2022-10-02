package cmd

import (
	"regexp"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hexdump",
	Short: "Display file contents in hexadecimal, decimal, octal, or ascii",
	Long:  "https://linux.die.net/man/1/hexdump",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("canonical", "C", false, "canonical hex+ASCII display")
	rootCmd.Flags().StringP("color", "L", "", "interpret color formatting specifiers")
	rootCmd.Flags().StringP("format", "e", "", "format string to be used for displaying data")
	rootCmd.Flags().StringP("format-file", "f", "", "file that contains format strings")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().StringP("length", "n", "", "interpret only length bytes of input")
	rootCmd.Flags().BoolP("no-squeezing", "v", false, "output identical lines")
	rootCmd.Flags().BoolP("one-byte-char", "c", false, "one-byte character display")
	rootCmd.Flags().BoolP("one-byte-octal", "b", false, "one-byte octal display")
	rootCmd.Flags().StringP("skip", "s", "", "skip offset bytes from the beginning")
	rootCmd.Flags().BoolP("two-bytes-decimal", "d", false, "two-byte decimal display")
	rootCmd.Flags().BoolP("two-bytes-hex", "x", false, "two-byte hexadecimal display")
	rootCmd.Flags().BoolP("two-bytes-octal", "o", false, "two-byte octal display")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("color").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":  carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"length": ActionMultiplicativeSuffixes(),
		"skip":   ActionMultiplicativeSuffixes(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}

func ActionMultiplicativeSuffixes() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		r := regexp.MustCompile(`^(?P<num>\d+)`)
		if r.MatchString(c.CallbackValue) {
			matches := r.FindStringSubmatch(c.CallbackValue)
			return carapace.ActionValues("KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB").Invoke(c).Prefix(matches[1]).ToA()
		}
		return carapace.ActionValues()
	})
}
