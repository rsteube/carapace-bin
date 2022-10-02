package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "basename",
	Short: "strip directory and suffix from filenames",
	Long:  "https://en.wikipedia.org/wiki/Basename",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("multiple", "a", false, "support multiple arguments and treat each as a NAME")
	rootCmd.Flags().StringP("suffix", "s", "", "remove a trailing SUFFIX; implies -a")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero", "z", false, "end each output line with NUL, not newline")
}
