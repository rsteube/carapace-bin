package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csplit",
	Short: "split a file into sections determined by context lines",
	Long:  "https://linux.die.net/man/1/csplit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("digits", "n", "", "use specified number of digits instead of 2")
	rootCmd.Flags().BoolP("elide-empty-files", "z", false, "remove empty output files")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("keep-files", "k", false, "do not remove output files on errors")
	rootCmd.Flags().StringP("prefix", "f", "", "use PREFIX instead of 'xx'")
	rootCmd.Flags().StringP("quiet", "s", "", "do not print counts of output file sizes")
	rootCmd.Flags().StringP("suffix-format", "b", "", "use sprintf FORMAT instead of %02d")
	rootCmd.Flags().Bool("suppress-matched", false, "suppress the lines matching PATTERN")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
