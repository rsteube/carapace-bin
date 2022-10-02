package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brotli",
	Short: "compress or decompress files",
	Long:  "https://github.com/google/brotli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteUnbrotli() error {
	rootCmd.Use = "unbrotli"
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "compression level")
	rootCmd.Flags().BoolS("1", "1", false, "compression level")
	rootCmd.Flags().BoolS("2", "2", false, "compression level")
	rootCmd.Flags().BoolS("3", "3", false, "compression level")
	rootCmd.Flags().BoolS("4", "4", false, "compression level")
	rootCmd.Flags().BoolS("5", "5", false, "compression level")
	rootCmd.Flags().BoolS("6", "6", false, "compression level")
	rootCmd.Flags().BoolS("7", "7", false, "compression level")
	rootCmd.Flags().BoolS("8", "8", false, "compression level")
	rootCmd.Flags().BoolS("9", "9", false, "compression level")
	rootCmd.Flags().BoolP("best", "Z", false, "use best compression level (11) (default)")
	rootCmd.Flags().BoolP("decompress", "d", false, "decompress")
	rootCmd.Flags().BoolP("force", "f", false, "force output file overwrite")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("keep", "k", false, "keep source file(s) (default)")
	rootCmd.Flags().String("large_window", "", "use incompatible large-window brotli")
	rootCmd.Flags().StringP("lgwin", "w", "", "set LZ77 window size")
	rootCmd.Flags().BoolP("no-copy-stat", "n", false, "do not copy source file(s) attributes")
	rootCmd.Flags().StringP("output", "o", "", "output file (only if 1 input file)")
	rootCmd.Flags().StringP("quality", "q", "", "compression level (0-11)")
	rootCmd.Flags().BoolP("rm", "j", false, "remove source file(s)")
	rootCmd.Flags().BoolP("stdout", "c", false, "write on standard output")
	rootCmd.Flags().StringP("suffix", "S", "", "output file suffix (default:'.br')")
	rootCmd.Flags().BoolP("test", "t", false, "test compressed file integrity")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	rootCmd.Flags().BoolP("version", "V", false, "display version and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output":  carapace.ActionFiles(),
		"quality": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"),
	})
	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
