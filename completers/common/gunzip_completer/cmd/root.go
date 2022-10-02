package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gunzip",
	Short: "Uncompress files",
	Long:  "https://linux.die.net/man/1/gunzip",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("force", "f", false, "force overwrite of output file and compress links")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("keep", "k", false, "keep (don't delete) input files")
	rootCmd.Flags().BoolP("list", "l", false, "list compressed file contents")
	rootCmd.Flags().BoolP("name", "N", false, "save or restore the original name and timestamp")
	rootCmd.Flags().BoolP("no-name", "n", false, "do not save or restore the original name and timestamp")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress all warnings")
	rootCmd.Flags().BoolP("recursive", "r", false, "operate recursively on directories")
	rootCmd.Flags().BoolP("stdout", "c", false, "write on standard output, keep original files unchanged")
	rootCmd.Flags().StringP("suffix", "S", "", "use suffix SUF on compressed files")
	rootCmd.Flags().BoolP("test", "t", false, "test compressed file integrity")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	rootCmd.Flags().Bool("version", false, "display version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
