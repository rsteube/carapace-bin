package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sed",
	Short: "stream editor for filtering and transforming text",
	Long:  "https://www.gnu.org/software/sed/manual/sed.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "use extended regular expressions in the script")
	rootCmd.Flags().Bool("debug", false, "annotate program execution")
	rootCmd.Flags().StringP("expression", "e", "", "add the script to the commands to be executed")
	rootCmd.Flags().StringP("file", "f", "", "add the contents of script-file to the commands to be executed")
	rootCmd.Flags().Bool("follow-symlinks", false, "follow symlinks when processing in place")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("in-place", "i", "", "edit files in place (makes backup if SUFFIX supplied)")
	rootCmd.Flags().StringP("line-length", "l", "", "specify the desired line-wrap length for the l command")
	rootCmd.Flags().BoolP("null-data", "z", false, "separate lines by NUL characters")
	rootCmd.Flags().Bool("posix", false, "disable all GNU extensions")
	rootCmd.Flags().StringP("quiet", "n", "", "suppress automatic printing of pattern space")
	rootCmd.Flags().BoolP("regexp-extended", "r", false, "use extended regular expressions in the script")
	rootCmd.Flags().Bool("sandbox", false, "operate in sandbox mode (disable e/r/w commands)")
	rootCmd.Flags().BoolP("seperate", "s", false, "consider files as separate rather than as a single, continuous long stream.")
	rootCmd.Flags().String("silent", "", "suppress automatic printing of pattern space")
	rootCmd.Flags().BoolP("unbuffered", "u", false, "load minimal amounts of data from the input files and flush the output buffers more often")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("expression").Changed || rootCmd.Flag("file").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
