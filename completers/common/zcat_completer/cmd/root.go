package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zcat",
	Short: "compress or expand files",
	Long:  "https://linux.die.net/man/1/zcat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "force; read compressed data even from a terminal")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("list", "l", false, "list compressed file contents")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress all warnings")
	rootCmd.Flags().BoolP("recursive", "r", false, "operate recursively on directories")
	rootCmd.Flags().StringP("suffix", "S", "", "use suffix SUF on compressed files")
	rootCmd.Flags().BoolP("test", "t", false, "test compressed file integrity")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	rootCmd.Flags().Bool("version", false, "display version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
