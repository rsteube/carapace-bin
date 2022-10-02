package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "du",
	Short: "estimate file space usage",
	Long:  "https://linux.die.net/man/1/du",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "equivalent to --dereference-args (-D)")
	rootCmd.Flags().BoolP("all", "a", false, "write counts for all files, not just directories")
	rootCmd.Flags().Bool("apparent-size", false, "print apparent sizes, rather than disk usage; although")
	rootCmd.Flags().StringP("block-size", "B", "", "scale sizes by SIZE before printing them; e.g.,")
	rootCmd.Flags().BoolP("bytes", "b", false, "equivalent to '--apparent-size --block-size=1'")
	rootCmd.Flags().BoolP("count-links", "l", false, "count sizes many times if hard linked")
	rootCmd.Flags().BoolP("dereference", "L", false, "dereference all symbolic links")
	rootCmd.Flags().BoolP("dereference-args", "D", false, "dereference only symlinks that are listed on the")
	rootCmd.Flags().String("exclude", "", "exclude files that match PATTERN")
	rootCmd.Flags().StringP("exclude-from", "X", "", "exclude files that match any pattern in FILE")
	rootCmd.Flags().String("files0-from", "", "summarize disk usage of the")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("human-readable", "h", false, "print sizes in human readable format (e.g., 1K 234M 2G)")
	rootCmd.Flags().Bool("inodes", false, "list inode usage information instead of block usage")
	rootCmd.Flags().BoolS("k", "k", false, "like --block-size=1K")
	rootCmd.Flags().BoolS("m", "m", false, "like --block-size=1M")
	rootCmd.Flags().StringP("max-depth", "d", "", "print the total for a directory (or file, with --all)")
	rootCmd.Flags().BoolP("no-dereference", "P", false, "don't follow any symbolic links (this is the default)")
	rootCmd.Flags().BoolP("null", "0", false, "end each output line with NUL, not newline")
	rootCmd.Flags().BoolP("one-file-system", "x", false, "skip directories on different file systems")
	rootCmd.Flags().BoolP("separate-dirs", "S", false, "for directories do not include size of subdirectories")
	rootCmd.Flags().Bool("si", false, "like -h, but use powers of 1000 not 1024")
	rootCmd.Flags().BoolP("summarize", "s", false, "display only a total for each argument")
	rootCmd.Flags().StringP("threshold", "t", "", "exclude entries smaller than SIZE if positive,")
	rootCmd.Flags().String("time", "", "show time as WORD instead of modification time:")
	rootCmd.Flags().String("time-style", "", "show times using STYLE, which can be:")
	rootCmd.Flags().BoolP("total", "c", false, "produce a grand total")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exclude-from": carapace.ActionFiles(),
		"files0-from":  carapace.ActionFiles(),
		"time":         carapace.ActionValues("access", "atime", "ctime", "status", "use"),
		"time-style":   carapace.ActionValues("full-iso", "long-iso", "iso"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
