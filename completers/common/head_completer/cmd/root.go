package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "head",
	Short: "output the first part of files",
	Long:  "https://linux.die.net/man/1/head",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bytes", "c", "", "print the first NUM bytes of each file;")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().StringP("lines", "n", "", "print the first NUM lines instead of the first 10;")
	rootCmd.Flags().BoolP("quiet", "q", false, "never print headers giving file names")
	rootCmd.Flags().Bool("silent", false, "never print headers giving file names")
	rootCmd.Flags().BoolP("verbose", "v", false, "always print headers giving file names")
	rootCmd.Flags().Bool("version", false, "output version information and exit")
	rootCmd.Flags().BoolP("zero-terminated", "z", false, "line delimiter is NUL, not newline")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
