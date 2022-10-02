package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xz",
	Short: "Compress or decompress .xz and .lzma files",
	Long:  "https://linux.die.net/man/1/xz",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func ExecuteAlias(alias string) error {
	rootCmd.Use = alias
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "compression preset")
	rootCmd.Flags().BoolS("1", "1", false, "compression preset")
	rootCmd.Flags().BoolS("2", "2", false, "compression preset")
	rootCmd.Flags().BoolS("3", "3", false, "compression preset")
	rootCmd.Flags().BoolS("4", "4", false, "compression preset")
	rootCmd.Flags().BoolS("5", "5", false, "compression preset")
	rootCmd.Flags().BoolS("6", "6", false, "compression preset")
	rootCmd.Flags().BoolS("7", "7", false, "compression preset")
	rootCmd.Flags().BoolS("8", "8", false, "compression preset")
	rootCmd.Flags().BoolS("9", "9", false, "compression preset")
	rootCmd.Flags().BoolP("compress", "z", false, "force compression")
	rootCmd.Flags().BoolP("decompress", "d", false, "force decompression")
	rootCmd.Flags().BoolP("extreme", "e", false, "try to improve compression ratio by using more CPU time;")
	rootCmd.Flags().BoolP("force", "f", false, "force overwrite of output file and (de)compress links")
	rootCmd.Flags().BoolP("help", "h", false, "display this short help and exit")
	rootCmd.Flags().BoolP("keep", "k", false, "keep (don't delete) input files")
	rootCmd.Flags().BoolP("list", "l", false, "list information about .xz files")
	rootCmd.Flags().BoolP("long-help", "H", false, "display the long help (lists also the advanced options)")
	rootCmd.Flags().BoolP("quiet", "q", false, "suppress warnings; specify twice to suppress errors too")
	rootCmd.Flags().BoolP("stdout", "c", false, "write to standard output and don't delete input files")
	rootCmd.Flags().BoolP("test", "t", false, "test compressed file integrity")
	rootCmd.Flags().StringP("threads", "T", "", "use at most NUM threads; the default is 1; set to 0")
	rootCmd.Flags().BoolP("verbose", "v", false, "be verbose; specify twice for even more verbose")
	rootCmd.Flags().BoolP("version", "V", false, "display the version number and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
